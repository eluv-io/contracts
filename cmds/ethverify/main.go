package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/eluv-io/contracts/cmds/ethverify/abidiff"
	"github.com/eluv-io/contracts/cmds/ethverify/gitutils"
)

const (
	DefaultElvFolder = ".eluvio/ethverify"
)

var (
	ostream log.Handler
	glogger = new(log.GlogHandler)
	config  string

	cmdRoot = &cobra.Command{
		Use:               "ethverify",
		Short:             "Management and verification of contracts",
		PersistentPreRunE: readConfig,
	}

	cmdGitFind = &cobra.Command{
		Use:   "git-find <contract_address> <Path/to/contracts/repo> <elvmasterd_rpc_url> ",
		Short: "Manage and retrieve the contract's git version",
		Long: `git-find helps to retrieve the git version at which the contract bytecode is present. 
The parameters can be set using flags or config file.`,
		Args: cobra.RangeArgs(2,3),
		RunE:    runGitFind,
		Example: `if running from contracts repo : ethverify git-find 0xCAFE . "http://localhost:8545"`,
	}

	cmdAbiDiff = &cobra.Command{
		Use:   "abi-diff",
		Short: "To identify changes made to contracts",
		Long: `abi-diff compares the new abi with stored abi and specifies any critical changes are present. 
If the changes made are not critical, the store dir has new set of abi. 
Also, the stored abi can be overwritten with new abi using --overwrite flag.
The default path for storedir is "./store", which can be changed using --storedir flag.
`,
		RunE: runAbiDiff,
		Example: `if run from 'contracts' repo, "ethverify abi-diff"
else "ethverify abi-diff --storedir <Path/to/stored/abi/dir>"
`,
	}
)

func init() {

	// cmd flags
	cmdRoot.PersistentFlags().Int("verbosity", 3, "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail")
	cmdRoot.PersistentFlags().String("log-file", "", "Output log file")
	cmdRoot.PersistentFlags().StringVar(&config, "config", "", "Config file path (default=<Homedir>/.eluvio/ethverify/config.toml")

	cmdAbiDiff.Flags().Bool("overwrite", false, "overwrite 'store' directory with new abi, even if abi-diff throws breaking changes")
	cmdAbiDiff.Flags().String("storedir", "./store", "directory for stored abi files")

	_ = viper.BindPFlag("verbosity", cmdRoot.PersistentFlags().Lookup("verbosity"))
	_ = viper.BindPFlag("log_file", cmdRoot.PersistentFlags().Lookup("log-file"))
	_ = viper.BindPFlag("git_find.ethurl", cmdGitFind.Flags().Lookup("ethurl"))
	_ = viper.BindPFlag("git_find.rootdir", cmdGitFind.Flags().Lookup("rootdir"))
	_ = viper.BindPFlag("git_find.contract", cmdGitFind.Flags().Lookup("contract"))
	_ = viper.BindPFlag("abi_diff.overwrite", cmdAbiDiff.Flags().Lookup("overwrite"))
	_ = viper.BindPFlag("abi_diff.storedir", cmdAbiDiff.Flags().Lookup("storedir"))

	// for env variable
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	cmdRoot.AddCommand(cmdGitFind)
	cmdRoot.AddCommand(cmdAbiDiff)

}

func readConfig(cmd *cobra.Command, args []string) error {

	// if --config is passed, attempt to parse the config file
	var filename string

	if config == "" {
		home, err := homedir.Dir()

		var folderPath string
		if err == nil {
			folderPath = path.Join(home, DefaultElvFolder)
			_, err := os.Stat(folderPath)
			if os.IsNotExist(err) {
				err = os.MkdirAll(folderPath, os.FileMode(0700))
				if err != nil {
					return fmt.Errorf("create config folder failed : folder = %v, err = %v", folderPath, err)
				}
			}
		}

		file := path.Join(folderPath, "config.toml")

		// checks if config file exists
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			log.Warn("reading from default config file", "filepath", file)
			filename = filepath.Base(file)
			viper.SetConfigName(filename[:len(filename)-len(filepath.Ext(filename))])
			viper.AddConfigPath(folderPath)

			err := viper.ReadInConfig()
			if err != nil {
				return fmt.Errorf("failed to read config file - %v", err)
			}
		}

	} else {
		filename = filepath.Base(config)
		viper.SetConfigName(filename[:len(filename)-len(filepath.Ext(filename))])
		viper.AddConfigPath(filepath.Dir(config))

		err := viper.ReadInConfig()
		if err != nil {
			return fmt.Errorf("failed to read config file - %v", err)
		}
	}

	var output io.Writer
	usecolor := (isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) && os.Getenv("TERM") != "dumb"
	logFile := viper.GetString("log_file")
	verbosity := viper.GetInt("verbosity")

	var lvl log.Lvl
	switch verbosity {
	case 0:
		lvl = log.LvlCrit
	case 1:
		lvl = log.LvlError
	case 2:
		lvl = log.LvlWarn
	case 3:
		lvl = log.LvlInfo
	case 4:
		lvl = log.LvlDebug
	case 5:
		lvl = log.LvlTrace
	}

	if usecolor {
		output = colorable.NewColorableStderr()
	}
	var err error
	if logFile != "" {
		ostream, err = log.FileHandler(logFile, log.TerminalFormat(false))
		if err != nil {
			utils.Fatalf("error setting logger file", err)
		}
	} else {
		output = io.Writer(os.Stdout)
		ostream = log.StreamHandler(output, log.TerminalFormat(usecolor))
	}

	glogger = log.NewGlogHandler(ostream)
	glogger.Verbosity(lvl)
	log.Root().SetHandler(glogger)

	return nil
}

func runGitFind(cmd *cobra.Command, args []string) error {

	contractAddr := args[0]
	if !common.IsHexAddress(contractAddr) || contractAddr == "" {
		return fmt.Errorf("contract address provided is invalid, contract addr = %v", contractAddr)
	}

	rootDir := args[1]
	if rootDir == "" {
		return fmt.Errorf("root directory is nil")
	}

	var ethurl string
	if len(args) > 2 {
		ethurl = args[2]
	} else {
		ethurl = "http://localhost:8545"
		log.Warn(fmt.Sprintf("elvmasterd_rpc_url:%v",ethurl))
	}

	gitCommits, err := gitutils.GetContractGitCommit(rootDir, ethurl, common.HexToAddress(contractAddr))
	if err != nil {
		return err
	}

	if len(gitCommits) > 0 {
		var table *tablewriter.Table
		table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Git Contract Info"})
		table.SetRowLine(true)
		for _, v := range gitCommits {
			table.Append([]string{v})
		}
		table.Render()
	} else{
		log.Info("No git commits are present for given contract address","contract address",contractAddr)
	}

	return nil
}

func runAbiDiff(cmd *cobra.Command, args []string) error {

	overwrite := viper.GetBool("abi_diff.overwrite")
	storeDir := viper.GetString("abi_diff.storedir")

	diffItem, err := abidiff.VerifyAllABI(overwrite, storeDir)
	if err != nil {
		return err
	}

	if len(diffItem) == 0 {
		log.Info("NO DIFFERENCES in abi")
		return nil
	}

	checkForBreakingChanges := false
	for _, v := range diffItem {
		if v.Breaking {
			checkForBreakingChanges = true
			log.Info("BREAKING Changes", "report", v.Report)
		} else {
			log.Info("Changes", "report", v.Report)
		}
	}

	// No Breaking changes, replace the stored abi
	if !checkForBreakingChanges {
		_, err = abidiff.VerifyAllABI(true, storeDir)
		if err != nil {
			return err
		}
		log.Info("No Breaking changes are made, replacing the STORED abi to NEW abi")
	}

	return nil
}

func Execute() error {
	err := cmdRoot.Execute()
	if err != nil {
		log.Error(err.Error())
		// seems post-run is not invoked in that case
		// cleanupServer(nil, nil)
	}
	return err
}

func main() {

	if err := Execute(); err != nil {
		os.Exit(1)
	}

	log.Info("LEAVING eth-verify...")
}
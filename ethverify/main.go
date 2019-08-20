package main

import (
	"fmt"
	"github.com/eluv-io/contracts/ethverify/gitutils"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	DefaultElvFolder = ".eluvio/ethverify"
)

var (
	ostream log.Handler
	glogger = new(log.GlogHandler)
	config  string

	cmdRoot = &cobra.Command{
		Use:   "ethverify",
		Short: "Manage and retrieve the contract's git version",
		Long: `This tool helps to retrieve the git version at which the contract bytecode is present. 
The parameters can be set using flags or config file.`,
		PersistentPreRunE: readConfig,
		RunE:              runEthVerify,
	}
)

func init() {

	// cmd flags
	cmdRoot.Flags().Int("verbosity", 3, "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail")
	cmdRoot.Flags().String("log-file", "", "Output log file")
	cmdRoot.Flags().StringVar(&config, "config", "", "Config file path (default=<Homedir>/.eluvio/ethverify/config.toml")
	cmdRoot.Flags().String("ethurl", "http://localhost:8545", "HTTP-RPC server listening interface")
	cmdRoot.Flags().String("rootdir", "", "git root directory")
	cmdRoot.Flags().String("contract", "", "provide contract address")

	_ = viper.BindPFlag("verbosity", cmdRoot.Flags().Lookup("verbosity"))
	_ = viper.BindPFlag("log_file", cmdRoot.Flags().Lookup("log-file"))
	_ = viper.BindPFlag("ethurl", cmdRoot.Flags().Lookup("ethurl"))
	_ = viper.BindPFlag("rootdir", cmdRoot.Flags().Lookup("rootdir"))
	_ = viper.BindPFlag("contract", cmdRoot.Flags().Lookup("contract"))

	// for env variable
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

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

func runEthVerify(cmd *cobra.Command, args []string) error {

	ethUrl := viper.GetString("ethurl")
	rootDir := viper.GetString("rootdir")
	if rootDir == "" {
		return fmt.Errorf("root directory is nil")
	}
	contractAddr := viper.GetString("contract")
	if !common.IsHexAddress(contractAddr) || contractAddr == "" {
		return fmt.Errorf("contract address provided is invalid, contract addr = %v", contractAddr)
	}
	gitCommits, err := gitutils.GetContractGitCommit(rootDir, ethUrl, common.HexToAddress(contractAddr))
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

	log.Info("EXECUTING eth-verify daemon...")

	if err := Execute(); err != nil {
		os.Exit(1)
	}

	log.Info("LEAVING eth-verify daemon...")
}

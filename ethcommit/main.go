package main

import (
	"fmt"
	"github.com/eluv-io/contracts/ethcommit/gitutils"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	DefaultElvFolder = ".eluvio/ethcommit"
)

var (
	ostream log.Handler
	glogger = new(log.GlogHandler)
	config  string

	cmdRoot = &cobra.Command{
		Use:               "ethcommit",
		Short:             "Manage and retrieve contract version",
		PersistentPreRunE: readConfig,
	}

	cmdGet = &cobra.Command{
		Use:   "get",
		Short: "get content version",
		RunE:  getCommit,
	}
)

func init() {

	// cmd flags
	cmdRoot.PersistentFlags().Int("verbosity", 3, "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail")
	cmdRoot.PersistentFlags().String("log-file", "", "Output log file")
	cmdRoot.PersistentFlags().StringVar(&config, "config", "", "Config file path")

	_ = viper.BindPFlag("node.verbosity", cmdRoot.PersistentFlags().Lookup("verbosity"))
	_ = viper.BindPFlag("node.log_file", cmdRoot.PersistentFlags().Lookup("log-file"))

	// for env variable
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	cmdRoot.AddCommand(cmdGet)

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
	logFile := viper.GetString("node.log_file")
	verbosity := viper.GetInt("node.verbosity")

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

func getCommit(cmd *cobra.Command, args []string) error {
	err := gitutils.GetContractGitCommit()
	if err != nil {
		return err
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

	log.Info("EXECUTING elv-master daemon...")

	if err := Execute(); err != nil {
		os.Exit(1)
	}

	log.Info("LEAVING elv-master daemon...")
}

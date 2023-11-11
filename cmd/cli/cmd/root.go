package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aocup",
	Short: "A wrapper around the aoc CLI https://github.com/scarvalhojr/aoc-cli",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.aocup.yaml)")

	defaultYear, _, defaultDay := time.Now().Date()
	rootCmd.PersistentFlags().IntP("year", "y", defaultYear, "event Year")
	rootCmd.PersistentFlags().IntP("day", "d", defaultDay, "event Day")
	rootCmd.PersistentFlags().StringP("output-path", "o", "", "path to output files to")

	err := viper.BindPFlags(rootCmd.PersistentFlags())
	checkError(err)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".aocup")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func checkError(err error) {
	if err != nil {
		slog.Default().Error(err.Error())
		os.Exit(1)
	}
}

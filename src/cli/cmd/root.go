package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type config struct {
	Year          int    `mapstructure:"year"`
	Day           int    `mapstructure:"day"`
	OutputPath    string `mapstructure:"output-path"`
	Overwrite     bool   `mapstructure:"overwrite"`
	Language      string `mapstructure:"language"`
	LeaderboardId string `mapstructure:"leaderboardId"`
	MemberId      string `mapstructure:"memberId"`
}

var conf config
var baseSolutionDirectory string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ao",
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

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.ao.yaml)")

	defaultYear, _, defaultDay := time.Now().Date()
	rootCmd.PersistentFlags().IntP("year", "y", defaultYear, "event year")
	rootCmd.PersistentFlags().IntP("day", "d", defaultDay, "event day")
	rootCmd.PersistentFlags().StringP("output-path", "o", "", "path to output files to (default is src/solutions/{year}/{day})")
	rootCmd.PersistentFlags().StringP("language", "l", "all", "specify writing templated files for a single language")

	err := viper.BindPFlags(rootCmd.PersistentFlags())
	cobra.CheckErr(err)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ao")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		slog.Default().Debug(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	} else {
		cobra.CheckErr(err)
	}
}

func setConfig() {
	err := viper.Unmarshal(&conf)
	cobra.CheckErr(err)

	if conf.OutputPath == "" {
		baseSolutionDirectory = "src/solutions"
		conf.OutputPath = path.Join(
			baseSolutionDirectory,
			fmt.Sprintf("%d/%02d", conf.Year, conf.Day),
		)
	}
}

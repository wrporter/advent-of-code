package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run all solutions for a year or a specified day",
	Run: func(cmd *cobra.Command, args []string) {
		if conf.Day == -1 {
			delimiter := ""
			for day := 1; day <= 25; day++ {
				filename := fmt.Sprintf("solutions/%d/%02d/go/main.go", conf.Year, day)
				if _, err := os.Stat(filename); err == nil {
					fmt.Print(delimiter)
					delimiter = "\n"
					_ = runDay(filename)
				}
			}
		} else {
			filename := fmt.Sprintf("solutions/%d/%02d/go/main.go", conf.Year, conf.Day)
			_ = runDay(filename)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		err := viper.BindPFlags(cmd.PersistentFlags())
		cobra.CheckErr(err)
		setConfig()
	},
}

func runDay(filename string) error {
	command := exec.Command("go", "run", filename)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().IntP("day", "d", -1, "Day to run. If set to -1, all days are run (default).")
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run solutions for a year and day (runs all days when day is 0)",
	Run: func(cmd *cobra.Command, args []string) {
		if conf.Day == 0 {
			delimiter := ""
			for day := 1; day <= 25; day++ {
				filename := path.Join(
					baseSolutionDirectory,
					fmt.Sprintf("%d/%02d", conf.Year, day),
					// TODO: Add multi-language support
					"go/main.go",
				)
				if _, err := os.Stat(filename); err == nil {
					fmt.Print(delimiter)
					delimiter = "\n"
					_ = runDay(filename)
				}
			}
		} else {
			filename := path.Join(conf.OutputPath, "go/main.go")
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
}

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wrporter/advent-of-code/cmd/ao/cmd/color"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "Record end time for a part",
	Run: func(cmd *cobra.Command, args []string) {
		path := conf.OutputPath
		if conf.Language != "all" {
			path = filepath.Join(conf.OutputPath, conf.Language)
		}
		path = filepath.Join(path, "time.json")

		part := viper.GetInt("part")

		file, err := os.ReadFile(path)
		cobra.CheckErr(err)

		var timings map[string]string
		err = json.Unmarshal(file, &timings)
		cobra.CheckErr(err)

		end := time.Now()
		timings[fmt.Sprintf("part%dEnd", part)] = end.Format(time.RFC3339)

		var start time.Time
		if part == 1 {
			start, err = time.Parse(time.RFC3339, timings["start"])
			cobra.CheckErr(err)
		} else {
			start, err = time.Parse(time.RFC3339, timings["part1End"])
			cobra.CheckErr(err)
		}

		elapsed := end.Sub(start).Round(time.Second)
		timings[fmt.Sprintf("part%dElapsed", part)] = fmt.Sprintf("%s", elapsed)

		content, err := json.MarshalIndent(timings, "", "  ")
		cobra.CheckErr(err)

		err = os.WriteFile(path, content, os.ModePerm)
		cobra.CheckErr(err)

		slog.Default().With("timings", timings).Info(fmt.Sprintf("⏱️ Times recorded at: %s", color.Set(color.Green, path)))
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		err := viper.BindPFlags(cmd.Flags())
		cobra.CheckErr(err)

		setConfig()
	},
}

func init() {
	rootCmd.AddCommand(endCmd)

	rootCmd.PersistentFlags().IntP("part", "p", 1, "part to record time for")
}

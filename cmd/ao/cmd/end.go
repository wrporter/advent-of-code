package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		checkError(err)

		var timings map[string]string
		err = json.Unmarshal(file, &timings)
		checkError(err)

		end := time.Now()
		timings[fmt.Sprintf("part%dEnd", part)] = end.Format(time.RFC3339)

		var start time.Time
		if part == 1 {
			start, err = time.Parse(time.RFC3339, timings["start"])
			checkError(err)
		} else {
			start, err = time.Parse(time.RFC3339, timings["part1End"])
			checkError(err)
		}

		elapsed := end.Sub(start).Round(time.Second)
		timings[fmt.Sprintf("part%dElapsed", part)] = fmt.Sprintf("%s", elapsed)

		content, err := json.MarshalIndent(timings, "", "  ")
		checkError(err)

		err = os.WriteFile(path, content, os.ModePerm)
		checkError(err)

		slog.Default().With("timings", timings).Info("⏱️ Times recorded!")
	},
}

func init() {
	cobra.OnInitialize(func() {
		err := viper.Unmarshal(&conf)
		checkError(err)
	})

	rootCmd.AddCommand(endCmd)

	endCmd.Flags().IntP("part", "p", 1, "part to record time for")

	err := viper.BindPFlags(endCmd.Flags())
	checkError(err)
}

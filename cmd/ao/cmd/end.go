package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wrporter/advent-of-code/cmd/ao/cmd/color"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "Record end time for a part. Download puzzle description after completing part 2.",
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

			downloadPuzzle()
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

func downloadPuzzle() {
	cmd := exec.Command(
		"aoc",
		"download",
		"--session-file",
		".adventofcode.session",
		"--year",
		fmt.Sprintf("%d", conf.Year),
		"--day",
		fmt.Sprintf("%d", conf.Day),
		"--overwrite",
		"--puzzle-file",
		filepath.Join(conf.OutputPath, "puzzle.md"),
		"--puzzle-only",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	cobra.CheckErr(err)

	slog.Default().Info("✅  Downloaded input")
}

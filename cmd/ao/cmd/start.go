package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/wrporter/advent-of-code/cmd/ao/cmd/color"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

type config struct {
	Year       int    `mapstructure:"year"`
	Day        int    `mapstructure:"day"`
	OutputPath string `mapstructure:"output-path"`
	Overwrite  bool   `mapstructure:"overwrite"`
	Language   string `mapstructure:"language"`
}

var conf config

func initConfigStart() {
	err := viper.Unmarshal(&conf)
	cobra.CheckErr(err)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start solving parts for a given Day",
	Run: func(cmd *cobra.Command, args []string) {
		generateNewDay()
		downloadInput()
		removeTrailingNewline()
		writeStartTime()
	},
}

func init() {
	cobra.OnInitialize(initConfigStart)
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Bool("overwrite", false, "overwrite existing files with templates")

	err := viper.BindPFlags(startCmd.Flags())
	cobra.CheckErr(err)
}

const (
	baseTemplateDir = "templates/0-template"
)

func generateNewDay() {
	templateExt := ".tmpl"
	var templatePaths []string
	if conf.Language == "all" {
		err := filepath.Walk(baseTemplateDir, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				templatePaths = append(templatePaths, path)
			}
			return nil
		})
		cobra.CheckErr(err)
	} else {
		paths, err := filepath.Glob(filepath.Join(baseTemplateDir, "*"+templateExt))
		cobra.CheckErr(err)
		templatePaths = append(templatePaths, paths...)

		err = filepath.Walk(filepath.Join(baseTemplateDir, conf.Language), func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				templatePaths = append(templatePaths, path)
			}
			return nil
		})
		cobra.CheckErr(err)
	}

	for _, templatePath := range templatePaths {
		newPath := strings.ReplaceAll(templatePath, baseTemplateDir+"/", "")
		newPath = strings.ReplaceAll(filepath.Join(conf.OutputPath, newPath), templateExt, "")
		err := os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
		cobra.CheckErr(err)

		if _, err := os.Stat(newPath); err == nil && !conf.Overwrite {
			slog.Default().Warn(fmt.Sprintf("Skipped writing file because it already exists: %s", color.Set(color.Yellow, newPath)))
			continue
		} else if err == nil {
			slog.Default().Warn(fmt.Sprintf("Overwriting file due to --overwrite flag: %s", color.Set(color.Yellow, newPath)))
		}

		file, err := os.Create(newPath)
		cobra.CheckErr(err)

		t, err := template.ParseFiles(templatePath)
		cobra.CheckErr(err)

		err = t.Execute(file, conf)
		cobra.CheckErr(err)

		slog.Default().Info(fmt.Sprintf("Wrote file: %s", color.Set(color.Green, newPath)))
	}

	slog.Default().Info("✅  Generated files")
}

func downloadInput() {
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
		"--input-file",
		filepath.Join(conf.OutputPath, "input.txt"),
		"--puzzle-file",
		filepath.Join(conf.OutputPath, "puzzle.md"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	cobra.CheckErr(err)

	slog.Default().Info("✅  Downloaded input")
}

func removeTrailingNewline() {
	inputPath := filepath.Join(conf.OutputPath, "input.txt")
	content, err := os.ReadFile(inputPath)
	cobra.CheckErr(err)

	err = os.WriteFile(inputPath, []byte(strings.TrimSuffix(string(content), "\n")), os.ModeType)
	cobra.CheckErr(err)

	slog.Default().Info("✅  Removed trailing newline from input")
}

func writeStartTime() {
	contents, err := json.MarshalIndent(map[string]string{
		"start": time.Now().Format(time.RFC3339),
	}, "", "  ")
	cobra.CheckErr(err)

	path := conf.OutputPath
	if conf.Language != "all" {
		path = filepath.Join(conf.OutputPath, conf.Language)
	}

	err = os.WriteFile(filepath.Join(path, "time.json"), contents, os.ModePerm)
	cobra.CheckErr(err)

	slog.Default().Info("✅  Wrote start time")
}

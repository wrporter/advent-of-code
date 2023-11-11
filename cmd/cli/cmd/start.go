package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
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
	checkError(err)

	if conf.OutputPath == "" {
		conf.OutputPath = fmt.Sprintf("solutions/%d/%02d", conf.Year, conf.Day)
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start solving parts for a given Day",
	Run: func(cmd *cobra.Command, args []string) {
		//slog.Default().With("config", conf).Info("Starting...")
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
	startCmd.Flags().StringP("output-path", "o", "", "path to output files to")
	startCmd.Flags().StringP("language", "l", "all", "specify writing templated files for a single language")

	err := viper.BindPFlags(startCmd.Flags())
	checkError(err)
}

const (
	baseTemplateDir = "templates/0-template"
)

func writeStartTime() {
	contents, err := json.MarshalIndent(map[string]string{
		"start": time.Now().Format(time.RFC3339),
	}, "", "  ")
	checkError(err)

	err = os.WriteFile(filepath.Join(conf.OutputPath, "time.json"), contents, os.ModePerm)
	checkError(err)

	slog.Default().Info("✅ Wrote start time")
}

func generateNewDay() {
	var templatePaths []string
	templateDir := baseTemplateDir
	if conf.Language != "all" {
		templateDir = filepath.Join(baseTemplateDir, conf.Language)
	}

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			templatePaths = append(templatePaths, path)
		}
		return nil
	})
	checkError(err)

	for _, templatePath := range templatePaths {
		newPath := strings.ReplaceAll(templatePath, baseTemplateDir+"/", "")
		newPath = strings.ReplaceAll(filepath.Join(conf.OutputPath, newPath), ".tmpl", "")
		err = os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
		checkError(err)

		if _, err := os.Stat(newPath); err == nil && !conf.Overwrite {
			slog.Default().With("file", newPath).Warn("Skipped writing file due to it already existing")
			continue
		} else if err == nil {
			slog.Default().With("file", newPath).Warn("Overwriting file due to overwrite flag")
		} else {
			slog.Default().With("file", newPath).Info("Writing file")
		}

		file, err := os.Create(newPath)
		checkError(err)

		t, err := template.ParseFiles(templatePath)
		checkError(err)

		err = t.Execute(file, conf)
		checkError(err)
	}

	slog.Default().Info("✅ Generated files")
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
	checkError(err)

	slog.Default().Info("✅ Downloaded input")
}

func removeTrailingNewline() {
	inputPath := filepath.Join(conf.OutputPath, "input.txt")
	content, err := os.ReadFile(inputPath)
	checkError(err)

	err = os.WriteFile(inputPath, []byte(strings.TrimSuffix(string(content), "\n")), os.ModeType)
	checkError(err)

	slog.Default().Info("✅ Removed trailing newline from input")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Data struct {
	Year int
	Day  int
}

var (
	baseTemplateDir = "templates/0-template"
	Reset           = "\033[0m"
	Red             = "\033[31m"
	Green           = "\033[32m"
	Yellow          = "\033[33m"
	Blue            = "\033[34m"
	Purple          = "\033[35m"
	Cyan            = "\033[36m"
	Gray            = "\033[37m"
	White           = "\033[97m"
)

type Config struct {
	Year       int
	Day        int
	OutputPath string
	Overwrite  bool
	Language   string
}

func main() {
	config := loadConfig()

	generateNewDay(config)
	log.Printf("✅ Generated files for %d/%d", config.Year, config.Day)

	downloadInput(config)
	log.Println("✅ Downloaded input")

	removeTrailingNewline(config)
	log.Println("✅ Removed trailing newline")
}

func loadConfig() *Config {
	defaultYear, _, defaultDay := time.Now().Date()
	year := flag.Int("year", defaultYear, "Event year.")
	day := flag.Int("day", defaultDay, "Event day.")
	overwrite := flag.Bool("overwrite", false, "Overwrite existing files.")
	language := flag.String("language", "", "Generate files only for a specific language.")
	flag.Parse()

	config := &Config{
		Year:       *year,
		Day:        *day,
		OutputPath: fmt.Sprintf("solutions/%d/%02d", *year, *day),
		Overwrite:  *overwrite,
		Language:   *language,
	}
	return config
}

func generateNewDay(config *Config) {
	var templatePaths []string
	templateDir := filepath.Join(baseTemplateDir, config.Language)
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
		newPath = strings.ReplaceAll(filepath.Join(config.OutputPath, newPath), ".tmpl", "")
		err = os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
		checkError(err)

		if _, err := os.Stat(newPath); err == nil && !config.Overwrite {
			log.Printf("%sSkipped writing file due to it already existing: %s%s", Yellow, newPath, Reset)
			continue
		} else if err == nil {
			log.Printf("%sOverwriting file due to overwrite flag: %s%s", Yellow, newPath, Reset)
		} else {
			log.Printf("%sWriting file: %s%s", Green, newPath, Reset)
		}

		file, err := os.Create(newPath)
		checkError(err)

		t, err := template.ParseFiles(templatePath)
		checkError(err)

		err = t.Execute(file, config)
		checkError(err)
	}
}

func downloadInput(config *Config) {

	//cmd := exec.Command(
	//	"aocdl",
	//	"-year",
	//	fmt.Sprintf("%d", year),
	//	"-day",
	//	fmt.Sprintf("%d", day),
	//	"-output",
	//	inputPath,
	//	"-overwrite",
	//)

	cmd := exec.Command(
		"aoc",
		"download",
		"--session-file",
		".adventofcode.session",
		"--year",
		fmt.Sprintf("%d", config.Year),
		"--day",
		fmt.Sprintf("%d", config.Day),
		"--overwrite",
		"--input-file",
		filepath.Join(config.OutputPath, "input.txt"),
		"--puzzle-file",
		filepath.Join(config.OutputPath, "puzzle.md"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	checkError(err)
}

func removeTrailingNewline(config *Config) {
	inputPath := filepath.Join(config.OutputPath, "input.txt")
	input, err := os.ReadFile(inputPath)
	checkError(err)
	err = os.WriteFile(inputPath, []byte(strings.TrimSuffix(string(input), "\n")), os.ModeType)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

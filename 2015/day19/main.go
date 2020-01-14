package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"regexp"
	"strings"
)

var replacementRegex = regexp.MustCompile(`^([a-zA-Z]+) => ([a-zA-Z]+)$`)

func parse(lines []string) (map[string][]string, string) {
	replacements := make(map[string][]string)
	for _, line := range lines[:len(lines)-2] {
		match := replacementRegex.FindStringSubmatch(line)
		replacements[match[1]] = append(replacements[match[1]], match[2])
	}
	return replacements, lines[len(lines)-1]
}

func combinations(replacements map[string][]string, medicineMolecule string) int {
	distinctMolecules := make(map[string]int)
	var regex *regexp.Regexp

	for oldValue, newValues := range replacements {
		for _, newValue := range newValues {
			regex = regexp.MustCompile(oldValue)
			indexes := regex.FindAllStringIndex(medicineMolecule, -1)
			for _, index := range indexes {
				if strings.Contains(medicineMolecule, oldValue) {
					replaced := replace(medicineMolecule, newValue, index[0], index[1])
					if _, ok := distinctMolecules[replaced]; ok {
						distinctMolecules[replaced]++
					} else {
						distinctMolecules[replaced] = 1
					}
				}
			}
		}
	}

	return len(distinctMolecules)
}

func replace(target string, value string, startIndex int, endIndex int) string {
	return target[:startIndex] + value + target[endIndex:]
}

func main() {
	lines, _ := file.ReadFile("./2015/day19/input.txt")
	fmt.Println(combinations(parse([]string{
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOHOHO",
	})))
	fmt.Println(combinations(parse(lines)))
}

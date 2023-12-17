package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	out.Day(2020, 4)
	input, _ := file.ReadFile("./2020/day4/input.txt")

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var requiredFields1 = map[string]func(string) bool{
	"byr": func(string) bool { return true },
	"iyr": func(string) bool { return true },
	"eyr": func(string) bool { return true },
	"hgt": func(string) bool { return true },
	"hcl": func(string) bool { return true },
	"ecl": func(string) bool { return true },
	"pid": func(string) bool { return true },
}

func part1(input []string) int {
	numValidPassports := 0

	for _, passport := range parsePassports(input) {
		if isValid(passport, requiredFields1) {
			numValidPassports++
		}
	}

	return numValidPassports
}

var heightRegex = regexp.MustCompile(`^(\d+)(cm|in)$`)
var hairColorRegex = regexp.MustCompile(`^#([0-9]|[a-f]){6}$`)
var eyeColorRegex = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var passportIDRegex = regexp.MustCompile(`^[0-9]{9}$`)

var requiredFields2 = map[string]func(value string) bool{
	"byr": func(value string) bool { return len(value) == 4 && inRange(value, 1920, 2002) },
	"iyr": func(value string) bool { return len(value) == 4 && inRange(value, 2010, 2020) },
	"eyr": func(value string) bool { return len(value) == 4 && inRange(value, 2020, 2030) },
	"hgt": func(value string) bool {
		if !heightRegex.MatchString(value) {
			return false
		}

		match := heightRegex.FindStringSubmatch(value)
		height := match[1]
		unit := match[2]

		return (unit == "cm" && inRange(height, 150, 193)) ||
			(unit == "in" && inRange(height, 59, 76))
	},
	"hcl": func(value string) bool { return hairColorRegex.MatchString(value) },
	"ecl": func(value string) bool { return eyeColorRegex.MatchString(value) },
	"pid": func(value string) bool { return passportIDRegex.MatchString(value) },
}

func part2(input []string) int {
	numValidPassports := 0

	passports := parsePassports(input)
	for _, passport := range passports {
		if isValid(passport, requiredFields2) {
			numValidPassports++
		}
	}

	return numValidPassports
}

func parsePassports(input []string) []map[string]string {
	var passports []map[string]string
	currentPassport := ""

	for i, line := range input {
		currentPassport += line + " "

		if line == "" || (i+1) == len(input) {
			currentPassport = strings.TrimSpace(currentPassport)
			parsedPassport := make(map[string]string)

			for _, field := range strings.Split(currentPassport, " ") {
				split := strings.Split(field, ":")
				parsedPassport[split[0]] = split[1]
			}

			passports = append(passports, parsedPassport)
			currentPassport = ""
		}
	}

	return passports
}

func isValid(passport map[string]string, requiredFields map[string]func(value string) bool) bool {
	for requiredField, isValidField := range requiredFields {
		fieldValue, ok := passport[requiredField]

		if !ok || !isValidField(fieldValue) {
			return false
		}
	}
	return true
}

func inRange(value string, min int, max int) bool {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	intValue := int(valueInt64)
	return err == nil && intValue >= min && intValue <= max
}

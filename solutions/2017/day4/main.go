package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 4
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	numValidPassphrases := 0

	for _, passphrase := range input {
		seen := make(map[string]bool)
		isValid := true

		for _, word := range strings.Fields(passphrase) {
			if seen[word] {
				isValid = false
				break
			}
			seen[word] = true
		}

		if isValid {
			numValidPassphrases++
		}
	}

	return numValidPassphrases
}

func part2(input []string) interface{} {
	numValidPassphrases := 0

	for _, passphrase := range input {
		isValid := true

		words := strings.Fields(passphrase)
		for i, word := range words {
			for j, word2 := range words {
				if i == j {
					continue
				}
				if mystrings.AreAnagram(word, word2) {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			numValidPassphrases++
		}
	}

	return numValidPassphrases
}

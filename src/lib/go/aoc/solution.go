package aoc

import (
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out/color"
	"aoc/src/lib/go/v2/timeit"
	"fmt"
	"os"
	"time"
)

// Solution is a runner for solving an Advent of Code puzzle for a given year
// and day. Some puzzles have multiple examples where an additional value is
// required as input. Each part executor function allows passing in additional
// arguments for such cases. This can also be useful in tests.
type Solution struct {
	Year int
	Day  int

	// Filename is the file that will be read as input and passed to each
	// respective part. Defaults to `input.txt`. Can be modified to point to a
	// different file, such as `sample.txt`.
	Filename string

	// Part1 is the executor to solve Part 1 of each AoC puzzle.
	Part1 func(input string, args ...interface{}) interface{}

	// Part2 is the executor to solve Part 1 of each AoC puzzle.
	Part2 func(input string, args ...interface{}) interface{}
}

// ReadInput reads the input file and encodes the contents into a string.
func (s Solution) ReadInput() string {
	filename := "input.txt"
	if len(s.Filename) > 0 {
		filename = s.Filename
	}

	bytes, _ := os.ReadFile(fmt.Sprintf("src/solutions/%d/%02d/%s", s.Year, s.Day, filename))
	return string(bytes)
}

// ReadInputFromTests is intended to be used by unit tests to read the input
// file. This is because tests are executed from the directory they reside in,
// whereas main functions execute from the context of the base directory of the
// project.
func (s Solution) ReadInputFromTests() string {
	bytes, _ := os.ReadFile("../input.txt")
	return string(bytes)
}

// Run is the primary runner and separates multiple parts of the processing.
//
//  1. Reads the input.
//  2. Times and executes the solution for Part 1.
//  3. Times and executes the solution for Part 2.
//  4. Prints output for both parts, along with the elapsed time it took for
//     each part to complete.
func (s Solution) Run(args1 []interface{}, args2 []interface{}) {
	fmt.Printf("🎄 %s%s%s%d: Day %d\n%s", color.Green, color.Underlined, color.Bold, s.Year, s.Day, color.Reset)

	input := s.ReadInput()

	start1 := time.Now()
	answer1 := s.Part1(input, args1...)
	elapsed1 := time.Since(start1)

	start2 := time.Now()
	answer2 := s.Part2(input, args2...)
	elapsed2 := time.Since(start2)
	elapsedTotal := time.Since(start1)

	padding := ints.Max(len(fmt.Sprintf("%v", answer1)), len(fmt.Sprintf("%v", answer2)))

	fmt.Printf("⭐  %sPart 1: %s%*v%s %s|%s 🕒 %s%s%s\n", color.Green, color.Red, -padding, answer1, color.Reset, color.Cyan, color.Reset, color.Purple, timeit.Round(elapsed1, 2), color.Reset)
	fmt.Printf("⭐  %sPart 2: %s%*v%s %s|%s 🕒 %s%s%s\n", color.Green, color.Red, -padding, answer2, color.Reset, color.Cyan, color.Reset, color.Purple, timeit.Round(elapsed2, 2), color.Reset)

	fmt.Printf("🕒 %sTotal: %s%s\n", color.Blue, timeit.Round(elapsedTotal, 2), color.Reset)
}

// ReadInputFile reads the given file, ignores any errors, and returns the
// contents encoded as a string.
func ReadInputFile(filename string) string {
	bytes, _ := os.ReadFile(filename)
	return string(bytes)
}

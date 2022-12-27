package solution

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/out/color"
	"github.com/wrporter/advent-of-code/internal/common/timeit2"
	"os"
	"time"
)

type Solution interface {
	Run(args1 []interface{}, args2 []interface{})
	Part1(input string, args ...interface{}) interface{}
	Part2(input string, args ...interface{}) interface{}
}

type AbstractSolution struct {
	Solution

	Year     int
	Day      int
	Filename string
}

func (s AbstractSolution) ReadInput() string {
	filename := "input.txt"
	if len(s.Filename) > 0 {
		filename = s.Filename
	}

	file := fmt.Sprintf("solutions/%d/%02d/%s", s.Year, s.Day, filename)
	bytes, _ := os.ReadFile(file)
	return string(bytes)
}

func (s AbstractSolution) Run(args1 []interface{}, args2 []interface{}) {
	fmt.Printf("🎄 %s%s%s%d: Day %d\n%s", color.Green, color.Underlined, color.Bold, s.Year, s.Day, color.Reset)

	input := s.ReadInput()

	start := time.Now()
	s.solvePart1(input, args1...)
	s.solvePart2(input, args2...)

	elapsed := time.Since(start)
	fmt.Printf("🕒 %sTotal: %s%s\n", color.Blue, timeit2.Round(elapsed, 2), color.Reset)
}

func (s AbstractSolution) solvePart1(input string, args ...interface{}) {
	start := time.Now()
	answer := s.Part1(input, args)
	elapsed := time.Since(start)

	fmt.Printf("⭐  %sPart 1: %s%v\n%s", color.Green, color.Red, answer, color.Reset)
	fmt.Printf("🕒 %s%s %s%s\n", color.Dim, "Part 1:", timeit2.Round(elapsed, 2), color.Reset)
}

func (s AbstractSolution) solvePart2(input string, args ...interface{}) {
	start := time.Now()
	answer := s.Part2(input, args)
	elapsed := time.Since(start)

	fmt.Printf("⭐  %sPart 2: %s%v\n%s", color.Green, color.Red, answer, color.Reset)
	fmt.Printf("🕒 %s%s %s%s\n", color.Dim, "Part 1:", timeit2.Round(elapsed, 2), color.Reset)
}

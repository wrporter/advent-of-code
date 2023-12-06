package solution

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out/color"
	"github.com/wrporter/advent-of-code/internal/common/v2/timeit"
	"os"
	"path"
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

func New(s Solution, year, day int) *AbstractSolution {
	return &AbstractSolution{
		Solution: s,
		Year:     year,
		Day:      day,
		Filename: "",
	}
}

func (s AbstractSolution) ReadInput() string {
	return s.ReadInputPrefix("")
}

func (s AbstractSolution) ReadInputPrefix(prefix string) string {
	filename := "input.txt"
	if len(s.Filename) > 0 {
		filename = s.Filename
	}

	file := fmt.Sprintf("solutions/%d/%02d/%s", s.Year, s.Day, filename)
	bytes, _ := os.ReadFile(path.Join(prefix, file))
	return string(bytes)
}

func (s AbstractSolution) Run(args1 []interface{}, args2 []interface{}) {
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

func (s AbstractSolution) solvePart1(input string, args ...interface{}) {
	start := time.Now()
	answer := s.Part1(input, args)
	elapsed := time.Since(start)

	fmt.Printf("⭐  %sPart 1: %s%-20v%s %s|%s 🕒 %s%s%s\n", color.Green, color.Red, answer, color.Reset, color.Cyan, color.Reset, color.Dim, timeit.Round(elapsed, 2), color.Reset)
}

func (s AbstractSolution) solvePart2(input string, args ...interface{}) {
	start := time.Now()
	answer := s.Part2(input, args)
	elapsed := time.Since(start)

	fmt.Printf("⭐  %sPart 2: %s%-20v%s %s|%s 🕒 %s%s%s\n", color.Green, color.Red, answer, color.Reset, color.Cyan, color.Reset, color.Dim, timeit.Round(elapsed, 2), color.Reset)
}

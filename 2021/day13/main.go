package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 13
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	dots, folds := parseInput(input)
	foldPaper(dots, folds[0])
	return len(dots)
}

func part2(input []string) interface{} {
	dots, folds := parseInput(input)
	for _, fold := range folds {
		foldPaper(dots, fold)
	}
	return "\n" + renderGrid(mapToGrid(dots))
}

func foldPaper(dots map[geometry.Point]bool, fold Fold) {
	for dot := range dots {
		if fold.axis == "x" {
			if dot.X == fold.position {
				delete(dots, dot)
			} else if dot.X > fold.position {
				delete(dots, dot)
				diff := ints.Abs(dot.X - fold.position)
				newDot := geometry.NewPoint(dot.X-(diff*2), dot.Y)
				dots[newDot] = true
			}
		}

		if fold.axis == "y" {
			if dot.Y == fold.position {
				delete(dots, dot)
			} else if dot.Y > fold.position {
				delete(dots, dot)
				diff := ints.Abs(dot.Y - fold.position)
				newDot := geometry.NewPoint(dot.X, dot.Y-(diff*2))
				dots[newDot] = true
			}
		}
	}
}

func parseInput(input []string) (map[geometry.Point]bool, []Fold) {
	dots := make(map[geometry.Point]bool)
	var folds []Fold
	onFolds := false
	for _, line := range input {
		if line == "" {
			onFolds = true
			continue
		}

		if onFolds {
			words := strings.Split(line, " ")
			parts := strings.Split(words[2], "=")
			fold := Fold{
				axis:     parts[0],
				position: convert.StringToInt(parts[1]),
			}
			folds = append(folds, fold)
		} else {
			parts, _ := convert.ToInts(strings.Split(line, ","))
			dot := geometry.NewPoint(parts[0], parts[1])
			dots[dot] = true
		}
	}

	return dots, folds
}

type Fold struct {
	axis     string
	position int
}

func mapToGrid(m map[geometry.Point]bool) (grid [][]string) {
	topLeft := geometry.Point{0, 0}
	bottomRight := geometry.Point{0, 0}
	for p := range m {
		topLeft.X = ints.Min(topLeft.X, p.X)
		topLeft.Y = ints.Max(topLeft.Y, p.Y)
		bottomRight.X = ints.Max(bottomRight.X, p.X)
		bottomRight.Y = ints.Min(bottomRight.Y, p.Y)
	}

	width := ints.Abs(topLeft.X) + ints.Abs(bottomRight.X) + 1
	height := ints.Abs(topLeft.Y) + ints.Abs(bottomRight.Y) + 1
	region := make([][]string, height)

	for y := 0; y < height; y++ {
		row := make([]string, width)
		region[y] = row
		my := topLeft.Y - y

		for x := 0; x < width; x++ {
			mx := topLeft.X + x
			if m[geometry.Point{mx, my}] {
				row[x] = "#"
			} else {
				row[x] = "."
			}
		}
	}

	return region
}

func displayGrid(grid [][]string) {
	out := &strings.Builder{}
	//out.WriteString("=====================================================\n")
	out.WriteString("\033c")
	out.WriteString(renderGrid(grid))
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func renderGrid(grid [][]string) string {
	result := ""
	for _, row := range grid {
		for _, spot := range row {
			result += spot
		}
		result += "\n"
	}
	return result
}

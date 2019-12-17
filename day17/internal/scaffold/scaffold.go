package scaffold

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day13/public/computer"
	"strings"
	"time"
)

type Scaffold struct {
	program *computer.Program
}

type Point struct {
	X int
	Y int
}

func New(code []int) *Scaffold {
	return &Scaffold{computer.NewProgram(code)}
}

func (s *Scaffold) SumAlignmentIntersections() int {
	cpu := computer.New()
	cpu.Run(s.program)
	done := false
	var grid [][]rune
	var row []rune

	for !done {
		if ascii, ok := <-s.program.Output; ok {
			if ascii == '\n' {
				grid = append(grid, row)
				row = make([]rune, 0)
			} else {
				row = append(row, rune(ascii))
			}
		} else {
			done = true
		}
	}

	intersections := getIntersections(grid)
	sum := 0
	for _, intersection := range intersections {
		sum += getAlignmentParameter(intersection)
	}

	displayGrid(grid)
	return sum
}

func getAlignmentParameter(intersection Point) int {
	return intersection.X * intersection.Y
}

func getIntersections(grid [][]rune) []Point {
	var intersections []Point
	for y, row := range grid {
		for x := range row {
			if isIntersection(grid, y, x) {
				intersections = append(intersections, Point{x, y})
			}
		}
	}
	return intersections
}

func isScaffold(grid [][]rune, y, x int) bool {
	return y >= 0 &&
		y < len(grid) &&
		x >= 0 &&
		x < len(grid[y]) &&
		grid[y][x] == '#'
}

func isIntersection(grid [][]rune, y, x int) bool {
	return isScaffold(grid, y, x) &&
		isScaffold(grid, y-1, x) &&
		isScaffold(grid, y+1, x) &&
		isScaffold(grid, y, x-1) &&
		isScaffold(grid, y, x+1)
}

func displayGrid(grid [][]rune) {
	out := &strings.Builder{}
	//out.WriteString("=====================================================\n")
	//out.WriteString("\033c")
	out.WriteString(renderGrid(grid))
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func renderGrid(grid [][]rune) string {
	result := ""
	for _, row := range grid {
		for _, spot := range row {
			result += string(spot)
		}
		result += "\n"
	}
	return result
}

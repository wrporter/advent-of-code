package solution

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/runegrid"
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"github.com/wrporter/advent-of-code/internal/common/v2/myslice"
	"os"
	"os/exec"
	"strings"
	"time"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	return countEnergized(grid, geometry.NewVector(0, 0, geometry.Right))
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	grid := convert.ToRuneGrid(strings.Split(input, "\n"))
	most := 0

	for y := range grid {
		most = max(most, countEnergized(grid, geometry.NewVector(0, y, geometry.Right)))
		most = max(most, countEnergized(grid, geometry.NewVector(len(grid[y])-1, y, geometry.Left)))
	}

	for x := range grid[0] {
		most = max(most, countEnergized(grid, geometry.NewVector(x, 0, geometry.Down)))
		most = max(most, countEnergized(grid, geometry.NewVector(x, len(grid)-1, geometry.Up)))
	}

	return most
}

func countEnergized(grid [][]rune, start *geometry.Vector) int {
	beams := []*geometry.Vector{start}
	seen := make(map[geometry.Vector]bool)

	for len(beams) > 0 {
		for _, beam := range beams {
			beams = beams[1:]

			x, y := beam.Point.X, beam.Point.Y
			if seen[*beam] || !myslice.InBounds(grid, y, x) {
				continue
			}
			seen[*beam] = true

			tile := grid[y][x]
			dir := beam.Direction

			if tile == '/' && (dir == geometry.Up || dir == geometry.Down) {
				beam.Rotate(90)
			} else if tile == '/' && (dir == geometry.Right || dir == geometry.Left) {
				beam.Rotate(-90)
			} else if tile == '\\' && (dir == geometry.Up || dir == geometry.Down) {
				beam.Rotate(-90)
			} else if tile == '\\' && (dir == geometry.Right || dir == geometry.Left) {
				beam.Rotate(90)
			}

			if tile == '|' && (dir == geometry.Right || dir == geometry.Left) {
				beams = append(beams, geometry.NewVector(x, y-1, geometry.Up))
				beams = append(beams, geometry.NewVector(x, y+1, geometry.Down))
			} else if tile == '-' && (dir == geometry.Up || dir == geometry.Down) {
				beams = append(beams, geometry.NewVector(x-1, y, geometry.Left))
				beams = append(beams, geometry.NewVector(x+1, y, geometry.Right))
			} else {
				beam.Move()
				beams = append(beams, beam)
			}

			//debug(grid, seen)
		}
	}

	energized := make(map[geometry.Point]struct{})
	for beam := range seen {
		energized[beam.Point] = struct{}{}
	}

	return len(energized)
}

func debug(grid [][]rune, seen map[geometry.Vector]bool) {
	next := myslice.Copy2D(grid)
	unique := make(map[geometry.Point]int)
	for beam := range seen {
		unique[beam.Point]++
	}

	for beam := range seen {
		x, y := beam.Point.X, beam.Point.Y
		if next[y][x] == '.' {
			if unique[beam.Point] > 1 {
				next[y][x] = '2'
			} else {
				next[y][x] = beam.Direction.Rune()
			}
		}
	}

	str := runegrid.String(next)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	fmt.Println(str)
	time.Sleep(100 * time.Millisecond)
}

package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"aoc/src/lib/go/ints"
	"bytes"
	"fmt"
	"strings"
	"time"
)

const (
	Bug        = '#'
	EmptySpace = '.'
)

type BugSensor struct {
	seen  map[string]int
	state [][]byte
}

func New(lines []string) *BugSensor {
	state := make([][]byte, len(lines))
	for y, line := range lines {
		row := make([]byte, len(line))
		for x, char := range line {
			row[x] = byte(char)
		}
		state[y] = row
	}

	seen := map[string]int{
		flatten(state): 1,
	}

	return &BugSensor{
		seen:  seen,
		state: state,
	}
}

func flatten(state [][]byte) string {
	var b bytes.Buffer
	for _, row := range state {
		b.Write(row)
	}
	return b.String()
}

func (b *BugSensor) FindMatchingLayout() int {
	for {
		b.Iterate()
		b.Display()
		if b.seen[flatten(b.state)] > 1 {
			return b.BiodiversityRating()
		}
	}
}

func (b *BugSensor) BiodiversityRating() int {
	flattened := flatten(b.state)
	rating := 0
	for i, value := range flattened {
		if value == Bug {
			rating += ints.Pow(2, i)
		}
	}
	return rating
}

func (b *BugSensor) Iterate() {
	next := make([][]byte, len(b.state))
	for y, row := range b.state {
		nextRow := make([]byte, len(row))
		for x := range row {
			nextRow[x] = EmptySpace
			if b.next(geometry.NewPoint(x, y)) {
				nextRow[x] = Bug
			}
		}
		next[y] = nextRow
	}
	b.seen[flatten(next)]++
	b.state = next
}

func (b *BugSensor) next(point geometry.Point) bool {
	isBug := b.state[point.Y][point.X] == Bug
	numAdjacentBugs := 0

	for _, direction := range geometry.Directions {
		p := point.Move(direction)
		if b.inBounds(p) && b.state[p.Y][p.X] == Bug {
			numAdjacentBugs++
		}
	}

	return (isBug && numAdjacentBugs == 1) ||
		(!isBug && (numAdjacentBugs == 1 || numAdjacentBugs == 2))
}

func (b *BugSensor) inBounds(p geometry.Point) bool {
	return p.Y >= 0 && p.Y < len(b.state) && p.X >= 0 && p.X < len(b.state[p.Y])
}

func (b *BugSensor) Display() {
	out := &strings.Builder{}
	out.WriteString("\033[2J\033[H")
	for _, row := range b.state {
		out.Write(row)
		out.WriteByte('\n')
	}
	fmt.Print(out.String())
	time.Sleep(time.Millisecond * 20)
}

func main() {
	lines, _ := file.ReadFile("./2019/day24/input.txt")
	sensor := New(lines)
	rating := sensor.FindMatchingLayout()
	fmt.Printf("Biodiversity rating: %d\n", rating)
}

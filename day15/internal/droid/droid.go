package droid

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day13/public/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/math"
	"strings"
	"time"
)

type Droid struct {
	program *computer.Program
	point   Point
}

type Point struct {
	X int
	Y int
}

func (p Point) North() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) South() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) West() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) East() Point {
	return Point{p.X + 1, p.Y}
}

func (p Point) Add(direction Direction) Point {
	x := p.X + directionModifiers[direction-1].X
	y := p.Y + directionModifiers[direction-1].Y
	return Point{x, y}
}

type Direction int

const (
	North Direction = 1
	South Direction = 2
	West  Direction = 3
	East  Direction = 4
)

var directions = []Direction{North, South, West, East}
var directionModifiers = []Point{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	default:
		return d
	}
}

type Status int

const (
	Wall         Status = 0
	Empty        Status = 1
	OxygenSystem Status = 2
)

func New(code []int) *Droid {
	return &Droid{computer.NewProgram(code), Point{0, 0}}
}

func (d *Droid) ScanShip() ([][]string, Point, Point) {
	cpu := computer.New()
	cpu.Run(d.program)
	m := map[Point]Status{Point{0, 0}: Empty}
	d.explore(m, Point{0, 0})
	//displayGrid(d.mapToGrid(m))
	return d.mapToGrid(m)
}

type node struct {
	point Point
	steps int
}

func (d *Droid) FindShortestPath(ship [][]string, from Point, to Point) int {
	visited := make(map[Point]bool)
	paths := make(map[Point]int)

	queue := []node{{from, 0}}
	var cur node

	for len(queue) > 0 {
		cur, queue = poll(queue)
		paths[cur.point] = cur.steps
		visited[cur.point] = true

		for _, direction := range directions {
			p := cur.point.Add(direction)
			if ship[p.Y][p.X] != WallSymbol && !visited[p] {
				queue = append(queue, node{p, cur.steps + 1})
			}
		}
	}

	return paths[to]
}

func (d *Droid) explore(m map[Point]Status, p Point) {
	for _, dir := range directions {
		next := p.Add(dir)
		if _, ok := m[next]; ok {
			continue
		}

		status := d.step(m, dir)
		if status == Wall {
			continue
		}

		d.explore(m, next)
		d.step(m, dir.Opposite())
	}
}

func (d *Droid) step(m map[Point]Status, dir Direction) Status {
	d.program.Input <- int(dir)
	status := Status(<-d.program.Output)

	next := d.point.Add(dir)
	if status == Empty || status == OxygenSystem {
		d.point = next
	}
	m[next] = status
	//displayGrid(d.mapToGrid(m))

	return status
}

func (d *Droid) mapToGrid(m map[Point]Status) (grid [][]string, start Point, oxygen Point) {
	topLeft := Point{0, 0}
	bottomRight := Point{0, 0}
	for p := range m {
		topLeft.X = math.Min(topLeft.X, p.X)
		topLeft.Y = math.Max(topLeft.Y, p.Y)
		bottomRight.X = math.Max(bottomRight.X, p.X)
		bottomRight.Y = math.Min(bottomRight.Y, p.Y)
	}

	width := math.Abs(topLeft.X) + math.Abs(bottomRight.X) + 1
	height := math.Abs(topLeft.Y) + math.Abs(bottomRight.Y) + 1
	region := make([][]string, height)

	for y := 0; y < height; y++ {
		row := make([]string, width)
		region[y] = row
		my := topLeft.Y - y

		for x := 0; x < width; x++ {
			mx := topLeft.X + x
			if spot, ok := m[Point{mx, my}]; ok {
				switch spot {
				case Wall:
					row[x] = WallSymbol
				case Empty:
					row[x] = EmptySymbol
				case OxygenSystem:
					oxygen = Point{x, y}
					row[x] = OxygenSystemSymbol
				default:
					row[x] = UnknownSymbol
				}
			} else {
				row[x] = UnknownSymbol
			}
		}
	}

	start = Point{-topLeft.X, topLeft.Y}
	region[topLeft.Y-d.point.Y][d.point.X-topLeft.X] = DroidSymbol
	region[start.Y][start.X] = StartSymbol

	return region, start, oxygen
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

const (
	UnknownSymbol      = " "
	WallSymbol         = "#"
	EmptySymbol        = "."
	OxygenSystemSymbol = "o"
	DroidSymbol        = "D"
	StartSymbol        = "X"
)

func poll(array []node) (node, []node) {
	return array[0], array[1:]
}

func pop(array []Direction) (Direction, []Direction) {
	size := len(array)
	return array[size-1], array[:size-1]
}

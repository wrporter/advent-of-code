package geometry

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/v2/mymath"
	"math"
	"strings"
)

type Direction int

const (
	Up    Direction = 1
	Right Direction = 2
	Down  Direction = 3
	Left  Direction = 4

	UpLeft    Direction = 5
	UpRight   Direction = 6
	DownLeft  Direction = 7
	DownRight Direction = 8
)

var Directions = []Direction{Up, Right, Down, Left}

var AllDirectionsModifiers = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
	{-1, -1},
	{1, -1},
	{-1, 1},
	{1, 1},
}

var DirectionModifiers = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

var DirectionModifiers2 = []Point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func (d Direction) Rotate(degrees int) Direction {
	return Directions[ints.WrapMod((int(d)-1)+(degrees*4/360), 4)]
}

type Vector struct {
	Point
	Direction
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Move(direction Direction) Point {
	x := p.X + AllDirectionsModifiers[direction-1].X
	y := p.Y + AllDirectionsModifiers[direction-1].Y
	return Point{X: x, Y: y}
}

func (p Point) Add(point Point) Point {
	x := p.X + point.X
	y := p.Y + point.Y
	return Point{X: x, Y: y}
}

func (p Point) AddAmount(direction Direction, amount int) Point {
	x := p.X + (DirectionModifiers2[direction-1].X * amount)
	y := p.Y + (DirectionModifiers2[direction-1].Y * amount)
	return Point{X: x, Y: y}
}

func (p Point) Rotate(degrees int) Point {
	degrees = ints.WrapMod(degrees, 360)
	if degrees == 90 {
		return Point{X: p.Y, Y: -p.X}
	} else if degrees == 180 {
		return Point{X: -p.X, Y: -p.Y}
	} else if degrees == 270 {
		return Point{X: -p.Y, Y: p.X}
	}
	return Point{X: p.X, Y: p.Y}
}

func (p Point) GetManhattanDistance() int {
	return ints.Abs(p.X) + ints.Abs(p.Y)
}

func (p Point) ManhattanDistance(p2 Point) int {
	return ints.Abs(p.X-p2.X) + ints.Abs(p.Y-p2.Y)
}

func (p Point) Up() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) Down() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) DownLeft() Point {
	return Point{X: p.X - 1, Y: p.Y + 1}
}

func (p Point) DownRight() Point {
	return Point{X: p.X + 1, Y: p.Y + 1}
}

func (p Point) Left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}

func (p Point) Right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func ToPoints(coordinateList []string) []Point {
	points := make([]Point, len(coordinateList))
	for i, coord := range coordinateList {
		points[i] = ToPoint(coord)
	}
	return points
}

func ToPoint(coordinates string) Point {
	coords, _ := convert.ToInts(strings.Split(coordinates, ","))
	return Point{X: coords[0], Y: coords[1]}
}

func MapToGrid(m map[Point]bool) []string {
	topLeft := NewPoint(math.MaxInt, math.MaxInt)
	bottomRight := NewPoint(math.MinInt, math.MinInt)
	for p := range m {
		topLeft.X = mymath.Min(topLeft.X, p.X)
		topLeft.Y = mymath.Min(topLeft.Y, p.Y)
		bottomRight.X = mymath.Max(bottomRight.X, p.X)
		bottomRight.Y = mymath.Max(bottomRight.Y, p.Y)
	}

	width := mymath.Abs(topLeft.X) + mymath.Abs(bottomRight.X) + 1
	height := mymath.Abs(topLeft.Y) + mymath.Abs(bottomRight.Y) + 1
	region := make([]string, height)

	for y := 0; y < height; y++ {
		dy := topLeft.Y + y

		for x := 0; x < width; x++ {
			dx := topLeft.X + x

			if m[NewPoint(dx, dy)] {
				region[y] += "#"
			} else {
				region[y] += "."
			}
		}
	}

	return region
}

func RenderGrid(grid []string) string {
	result := ""
	for _, row := range grid {
		result += row
		result += "\n"
	}
	return result
}

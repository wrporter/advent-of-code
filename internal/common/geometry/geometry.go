package geometry

import "github.com/wrporter/advent-of-code/internal/common/ints"

type Direction int

const (
	Up    Direction = 1
	Right Direction = 2
	Down  Direction = 3
	Left  Direction = 4
)

var Directions = []Direction{Up, Right, Down, Left}

var AllDirections = []Point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
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
	return Point{x, y}
}

func (p Point) Move(direction Direction) Point {
	x := p.X + DirectionModifiers[direction-1].X
	y := p.Y + DirectionModifiers[direction-1].Y
	return Point{x, y}
}

func (p Point) Add(point Point) Point {
	x := p.X + point.X
	y := p.Y + point.Y
	return Point{x, y}
}

func (p Point) AddAmount(direction Direction, amount int) Point {
	x := p.X + (DirectionModifiers2[direction-1].X * amount)
	y := p.Y + (DirectionModifiers2[direction-1].Y * amount)
	return Point{x, y}
}

func (p Point) Rotate(degrees int) Point {
	degrees = ints.WrapMod(degrees, 360)
	if degrees == 90 {
		return Point{p.Y, -p.X}
	} else if degrees == 180 {
		return Point{-p.X, -p.Y}
	} else if degrees == 270 {
		return Point{-p.Y, p.X}
	}
	return Point{p.X, p.Y}
}

func (p Point) GetManhattanDistance() int {
	return ints.Abs(p.X) + ints.Abs(p.Y)
}

func (p Point) Up() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) Left() Point {
	return Point{p.X - 1, p.Y}
}

func (p Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

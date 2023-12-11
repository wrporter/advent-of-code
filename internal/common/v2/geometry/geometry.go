package geometry

import (
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/v2/mymath"
	"math"
	"strconv"
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

func (d Direction) Rotate(degrees int) Direction {
	return Directions[ints.WrapMod((int(d)-1)+(degrees*4/360), 4)]
}

type Vector struct {
	*Point
	Direction
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) Copy() *Point {
	return NewPoint(p.X, p.Y)
}

func (p *Point) Move(direction Direction) *Point {
	p.X += AllDirectionsModifiers[direction-1].X
	p.Y += AllDirectionsModifiers[direction-1].Y
	return p
}

func (p *Point) Add(point *Point) {
	p.X += point.X
	p.Y += point.Y
}

func (p *Point) MoveAmount(direction Direction, amount int) {
	p.X += AllDirectionsModifiers[direction-1].X * amount
	p.Y += AllDirectionsModifiers[direction-1].Y * amount
}

func (p *Point) Rotate(degrees int) {
	degrees = ints.WrapMod(degrees, 360)
	if degrees == 90 {
		p.X = p.Y
		p.Y = -p.X
	} else if degrees == 180 {
		p.X = -p.X
		p.Y = -p.Y
	} else if degrees == 270 {
		p.X = -p.Y
		p.Y = p.X
	}
}

func (p *Point) GetManhattanDistance() int {
	return ints.Abs(p.X) + ints.Abs(p.Y)
}

func (p *Point) ManhattanDistance(p2 Point) int {
	return ints.Abs(p.X-p2.X) + ints.Abs(p.Y-p2.Y)
}

func (p *Point) Distance(p2 *Point) int {
	return mymath.Max(mymath.Abs(p.X-p2.X), mymath.Abs(p.Y-p2.Y))
}

func (p *Point) Diff(p2 *Point) (int, int) {
	return p.X - p2.X, p.Y - p2.Y
}

func (p *Point) Equals(p2 *Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *Point) Up() {
	p.Y -= 1
}

func (p *Point) Down() {
	p.Y += 1
}

func (p *Point) Left() {
	p.X -= 1
}

func (p *Point) Right() {
	p.X += 1
}

func (p *Point) DownLeft() {
	p.X -= 1
	p.Y += 1
}

func (p *Point) DownRight() {
	p.X += 1
	p.Y += 1
}

func (p *Point) UpLeft() {
	p.X -= 1
	p.Y -= 1
}

func (p *Point) UpRight() {
	p.X += 1
	p.Y -= 1
}

func (p *Point) String() string {
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(p.X))
	builder.WriteString(",")
	builder.WriteString(strconv.Itoa(p.Y))
	return builder.String()
}

func ToPoints(coordinateList []string) []*Point {
	points := make([]*Point, len(coordinateList))
	for i, coord := range coordinateList {
		points[i] = ToPoint(coord)
	}
	return points
}

func ToPoint(coordinates string) *Point {
	coords, _ := convert.ToInts(strings.Split(coordinates, ","))
	return &Point{X: coords[0], Y: coords[1]}
}

type GridMap struct {
	Grid [][]uint8
	MinY int
	MinX int
}

func MapToGridV2[T string | bool](m map[Point]T) *GridMap {
	topLeft := NewPoint(math.MaxInt, math.MaxInt)
	bottomRight := NewPoint(math.MinInt, math.MinInt)
	for p := range m {
		topLeft.X = mymath.Min(topLeft.X, p.X)
		topLeft.Y = mymath.Min(topLeft.Y, p.Y)
		bottomRight.X = mymath.Max(bottomRight.X, p.X)
		bottomRight.Y = mymath.Max(bottomRight.Y, p.Y)
	}

	width := mymath.Abs(topLeft.X-bottomRight.X) + 1
	height := mymath.Abs(topLeft.Y-bottomRight.Y) + 1
	grid := make([][]uint8, height)

	for y := 0; y < height; y++ {
		grid[y] = make([]uint8, width)
		dy := topLeft.Y + y

		for x := 0; x < width; x++ {
			dx := topLeft.X + x
			value, exists := m[*NewPoint(dx, dy)]
			grid[y][x] = getChar(value, exists)
		}
	}

	return &GridMap{
		Grid: grid,
		MinY: topLeft.Y,
		MinX: topLeft.X,
	}
}

func (g *GridMap) Translate(p *Point) *Point {
	return NewPoint(p.X-g.MinX, p.Y-g.MinY)
}

func Imprint[T string | bool](g *GridMap, m map[Point]T) {
	for p, v := range m {
		y := p.Y - g.MinY
		x := p.X - g.MinX
		g.Grid[y][x] = getChar(v, true)
	}
}

func getChar[T string | bool](value T, exists bool) uint8 {
	if !exists {
		return '.'
	}

	switch any(value).(type) {
	case string:
		return any(value).(string)[0]
	case bool:
		return '#'
	}
	return '.'
}

func MapToGrid[T string | bool](m map[Point]T) []string {
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

			if value, ok := m[*NewPoint(dx, dy)]; ok {
				switch any(value).(type) {
				case string:
					region[y] += any(value).(string)
				case bool:
					region[y] += "#"
				}
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

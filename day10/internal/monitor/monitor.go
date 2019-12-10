package monitor

import (
	"github.com/wrporter/advent-of-code-2019/internal/common/math"
	"strings"
)

type Monitor struct{}
type Station struct {
	NumVisibleAsteroids int
	Point               Point
}
type Point struct {
	X float64
	Y float64
}

func NewPoint(x int, y int) Point {
	return Point{float64(x), float64(y)}
}

func (p Point) slope(p1 Point) float64 {
	return (p.Y - p1.Y) / (p.X - p.X)
}
func (p Point) equals(p1 Point) bool {
	return p.Y == p1.Y && p.X == p1.X
}

type Line struct {
	P1 Point
	P2 Point
}

func (l Line) slope() float64 {
	return (l.P2.Y - l.P1.Y) / (l.P2.X - l.P1.X)
}
func (l Line) slopeIsUndefined() bool {
	return l.P2.X-l.P1.X == 0
}
func (l Line) intercept() float64 {
	return l.P1.Y - l.slope()*l.P1.X
}
func (l Line) contains(p Point) bool {
	if l.slopeIsUndefined() {
		return p.X == l.P1.X && p.X == l.P2.X
	}
	return p.Y == math.ToFixed(l.slope()*p.X+l.intercept(), 2)
}

const (
	Empty    = "."
	Asteroid = "#"
)

func New() *Monitor {
	return &Monitor{}
}

func (m *Monitor) FindBestAsteroid(field [][]string) Station {
	station := Station{}

	for y, row := range field {
		for x, thing := range row {
			if thing == Asteroid {
				p := NewPoint(x, y)
				count := countVisibleAsteroids(field, p)

				if count > station.NumVisibleAsteroids {
					station.NumVisibleAsteroids = count
					station.Point = p
				}
			}
		}
	}

	return station
}

func countVisibleAsteroids(field [][]string, station Point) int {
	count := 0

	for y, row := range field {
		for x, thing := range row {
			p := NewPoint(x, y)

			if thing == Asteroid &&
				!station.equals(p) &&
				canSeeAsteroid(field, station, p) {
				count++
			}
		}
	}

	return count
}

func canSeeAsteroid(field [][]string, station Point, asteroid Point) bool {
	line := Line{station, asteroid}

	for y := math.Min(int(station.Y), int(asteroid.Y)); y <= math.Max(int(station.Y), int(asteroid.Y)); y++ {
		for x := math.Min(int(station.X), int(asteroid.X)); x <= math.Max(int(station.X), int(asteroid.X)); x++ {
			p := NewPoint(x, y)

			if field[y][x] == Asteroid &&
				!station.equals(p) &&
				!asteroid.equals(p) &&
				line.contains(p) {
				return false
			}
		}
	}

	return true
}

func SplitLines(lines []string) [][]string {
	result := make([][]string, len(lines))
	for i, line := range lines {
		result[i] = strings.Split(line, "")
	}
	return result
}

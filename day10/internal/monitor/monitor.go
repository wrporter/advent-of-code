package monitor

import (
	"math"
	"sort"
	"strings"
)

type Monitor struct{}
type Station struct {
	NumVisibleAsteroids int
	Point               Point
}
type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func (p Point) slope(p1 Point) float64 {
	return float64((p.Y - p1.Y) / (p.X - p.X))
}
func (p Point) equals(p1 Point) bool {
	return p.Y == p1.Y && p.X == p1.X
}
func (p Point) distance(p1 Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-p1.X), 2) + math.Pow(float64(p.Y-p1.Y), 2))
}

func Angle(origin Point, target Point) float64 {
	degrees := 180 - (math.Atan2(float64(origin.Y-target.Y), float64(origin.X-target.X)))*180/math.Pi
	if degrees < 0 {
		degrees += 360
	}
	return degrees
}

const (
	Empty    = "."
	Asteroid = "#"
)

func New() *Monitor {
	return &Monitor{}
}

func (m *Monitor) FindBestAsteroid(field [][]string) (station Station, asteroids map[float64][]Point) {
	for y, row := range field {
		for x, thing := range row {
			if thing == Asteroid {
				p := NewPoint(x, y)
				candidate := getVisibleAsteroids(field, p)

				if len(candidate) > station.NumVisibleAsteroids {
					asteroids = candidate
					station.NumVisibleAsteroids = len(candidate)
					station.Point = p
				}
			}
		}
	}

	return station, asteroids
}

func (m *Monitor) ZapAsteroids(field [][]string, target int) (int, *Point) {
	station, asteroids := m.FindBestAsteroid(field)
	angles := clockwise(asteroids)

	for _, points := range asteroids {
		sort.Slice(points, func(i, j int) bool {
			return points[i].distance(station.Point) < points[j].distance(station.Point)
		})
	}

	numZapped := 0
	var asteroid Point

	for numZapped < target {
		for _, angle := range angles {
			if len(asteroids[angle]) > 0 {
				numZapped++
				asteroid, asteroids[angle] = poll(asteroids[angle])
				if numZapped == target {
					return asteroid.X*100 + asteroid.Y, &asteroid
				}
			}
		}
	}

	return 0, nil
}

func getVisibleAsteroids(field [][]string, station Point) map[float64][]Point {
	asteroids := make(map[float64][]Point)

	for y, row := range field {
		for x, thing := range row {
			asteroid := NewPoint(x, y)

			if thing == Asteroid && !station.equals(asteroid) {
				angle := Angle(station, asteroid)
				asteroids[angle] = append(asteroids[angle], asteroid)
			}
		}
	}

	return asteroids
}

func clockwise(m map[float64][]Point) []float64 {
	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return shift(keys[i]) > shift(keys[j])
	})
	return keys
}
func shift(angle float64) float64 {
	return math.Mod(angle+270, 361)
}

func SplitLines(lines []string) [][]string {
	result := make([][]string, len(lines))
	for i, line := range lines {
		result[i] = strings.Split(line, "")
	}
	return result
}

func poll(array []Point) (Point, []Point) {
	return array[0], array[1:]
}

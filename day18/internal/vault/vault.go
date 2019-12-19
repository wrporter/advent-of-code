package vault

import (
	"github.com/wrporter/advent-of-code-2019/internal/common/arrays"
	"github.com/wrporter/advent-of-code-2019/internal/common/bytes"
	"github.com/wrporter/advent-of-code-2019/internal/common/geometry"
	"github.com/wrporter/advent-of-code-2019/internal/common/mystrings"
	"strings"
)

const (
	StoneWall   = '#'
	Entrance    = '@'
	OpenPassage = '.'
)

type Vault struct {
	maze     map[geometry.Point]byte
	entrance geometry.Point
	keys     map[byte]geometry.Point
	doors    map[byte]geometry.Point
}

func New(vault []string) *Vault {
	maze := make(map[geometry.Point]byte)
	var entrance geometry.Point
	keys := make(map[byte]geometry.Point)
	doors := make(map[byte]geometry.Point)

	for y, line := range vault {
		for x, char := range line {
			spot := byte(char)

			if spot != StoneWall {
				maze[geometry.NewPoint(x, y)] = spot
			}
			if spot == Entrance {
				entrance = geometry.NewPoint(x, y)
			}
			if isKey(spot) {
				keys[spot] = geometry.NewPoint(x, y)
			}
			if isDoor(spot) {
				doors[spot] = geometry.NewPoint(x, y)
			}
		}
	}

	return &Vault{
		maze:     maze,
		entrance: entrance,
		keys:     keys,
		doors:    doors,
	}
}

type VisitedIteration struct {
	point       geometry.Point
	currentKeys string
}

func (v *Vault) MinSteps() int {
	return v.minSteps(v.entrance, "", make(map[VisitedIteration]int))
}

func (v *Vault) minSteps(from geometry.Point, currentKeys string, visited map[VisitedIteration]int) int {
	iteration := VisitedIteration{from, currentKeys}
	if steps, ok := visited[iteration]; ok {
		return steps
	}

	steps := 0
	keys := v.findKeys(from, currentKeys)

	if len(keys) > 0 {
		var possibleSteps []int
		for key, distance := range keys {
			cur := distance + v.minSteps(v.keys[key], mystrings.SortString(currentKeys+string(key)), visited)
			possibleSteps = append(possibleSteps, cur)
		}
		steps = arrays.Min(possibleSteps)
	}

	visited[iteration] = steps
	return steps
}

type Node struct {
	Element byte
	Point   geometry.Point
}

func (v *Vault) findKeys(from geometry.Point, currentKeys string) map[byte]int {
	keys := make(map[byte]int)
	distance := map[geometry.Point]int{from: 0}
	queue := []geometry.Point{from}
	var current geometry.Point

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, direction := range geometry.Directions {
			point := current.Add(direction)
			if _, ok := distance[point]; ok {
				continue
			}

			if item, ok := v.maze[point]; ok {
				distance[point] = distance[current] + 1

				if isDoor(item) && !hasKey(currentKeys, getKey(item)) {
					continue
				}
				if isKey(item) && !hasKey(currentKeys, item) {
					keys[item] = distance[point]
				} else {
					queue = append(queue, point)
				}
			}
		}
	}

	return keys
}

func hasKey(keys string, key byte) bool {
	return strings.ContainsRune(keys, rune(key))
}

func getKey(door byte) byte {
	return bytes.ToLower(door)
}

func getDoor(key byte) byte {
	return bytes.ToUpper(key)
}

func isDoor(b byte) bool {
	return bytes.IsUpper(b)
}

func isKey(b byte) bool {
	return bytes.IsLower(b)
}

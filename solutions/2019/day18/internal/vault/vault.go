package vault

import (
	"github.com/wrporter/advent-of-code/internal/common/arrays"
	"github.com/wrporter/advent-of-code/internal/common/bytes"
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"strings"
)

const (
	StoneWall   = '#'
	Entrance    = '@'
	OpenPassage = '.'
)

type Vault struct {
	maze      map[geometry.Point]byte
	entrances []geometry.Point
	keys      map[byte]geometry.Point
	doors     map[byte]geometry.Point
	distances map[byte][]Distance
}

func New(vault []string) *Vault {
	maze := make(map[geometry.Point]byte)
	keys := make(map[byte]geometry.Point)
	var entrances []geometry.Point
	doors := make(map[byte]geometry.Point)
	var entranceCounter byte = 0

	for y, line := range vault {
		for x, char := range line {
			spot := byte(char)

			if spot != StoneWall {
				maze[geometry.NewPoint(x, y)] = spot
			}
			if spot == Entrance {
				entrance := geometry.NewPoint(x, y)
				keys[entranceCounter] = entrance
				entranceCounter++
				entrances = append(entrances, entrance)
			}
			if isKey(spot) {
				keys[spot] = geometry.NewPoint(x, y)
			}
			if isDoor(spot) {
				doors[spot] = geometry.NewPoint(x, y)
			}
		}
	}

	// Pruning does not actually speed up performance because the BFS search does not go far into dead ends.
	//pruneDeadEnds(maze)
	distances := computeDistances(maze, keys)

	return &Vault{
		maze:      maze,
		keys:      keys,
		doors:     doors,
		distances: distances,
		entrances: entrances,
	}
}

type CacheKey struct {
	from        string
	currentKeys string
}

func (v *Vault) MinSteps() int {
	entranceKeys := make([]byte, len(v.entrances))
	for i := range v.entrances {
		entranceKeys[i] = byte(i)
	}
	return v.minSteps(entranceKeys, "", make(map[CacheKey]int))
}

func (v *Vault) minSteps(from []byte, currentKeys string, cache map[CacheKey]int) int {
	iteration := CacheKey{string(from), currentKeys}
	if steps, ok := cache[iteration]; ok {
		return steps
	}

	steps := 0
	keys := v.getReachableKeys(from, currentKeys)

	if len(keys) > 0 {
		var possibleSteps []int
		for _, distance := range keys {
			nextFrom := bytes.Copy(from)
			nextFrom[distance.robot] = distance.key
			cur := distance.distance + v.minSteps(nextFrom, mystrings.SortString(currentKeys+string(distance.key)), cache)
			possibleSteps = append(possibleSteps, cur)
		}
		steps = arrays.Min(possibleSteps)
	}

	cache[iteration] = steps
	return steps
}

type KeyDistance struct {
	robot    byte
	key      byte
	distance int
}

func (v *Vault) getReachableKeys(from []byte, currentKeys string) []KeyDistance {
	var keys []KeyDistance
	for robot, fromKey := range from {
		for _, distance := range v.distances[fromKey] {
			if !hasKey(currentKeys, distance.key) &&
				len(strings.Trim(string(distance.neededKeys), currentKeys)) == 0 {
				keys = append(keys, KeyDistance{byte(robot), distance.key, distance.distance})
			}
		}
	}
	return keys
}

//	'@': {
//			key: 'b',
//			neededKeys: {'A', 'B'}
//			distance: 5
//	},
//
// 'a': { ... }, ...
type Distance struct {
	key        byte
	neededKeys []byte
	distance   int
}

type node struct {
	point      geometry.Point
	neededKeys []byte
}

func computeDistances(maze map[geometry.Point]byte, startKeys map[byte]geometry.Point) map[byte][]Distance {
	distances := make(map[byte][]Distance)

	for key, point := range startKeys {
		var keys []Distance
		distance := map[geometry.Point]int{point: 0}
		queue := []node{{point, nil}}
		var current node

		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]

			for _, direction := range geometry.Directions {
				nextPoint := current.point.Move(direction)
				if _, ok := distance[nextPoint]; ok {
					continue
				}

				if tile, ok := maze[nextPoint]; ok {
					distance[nextPoint] = distance[current.point] + 1

					if isKey(tile) {
						keys = append(keys, Distance{tile, current.neededKeys, distance[nextPoint]})
					}

					if isDoor(tile) {
						queue = append(queue, node{nextPoint, bytes.CopyAdd(current.neededKeys, getKey(tile))})
					} else {
						queue = append(queue, node{nextPoint, current.neededKeys})
					}
				}
			}
		}
		distances[key] = keys
	}
	return distances
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

func copyAddMap(m map[byte]geometry.Point, key byte, value geometry.Point) map[byte]geometry.Point {
	target := make(map[byte]geometry.Point)
	for key, value := range m {
		target[key] = value
	}
	target[key] = value
	return target
}

func pruneDeadEnds(maze map[geometry.Point]byte) {
	for point := range maze {
		fillDeadEnd(maze, point)
	}
}

func fillDeadEnd(maze map[geometry.Point]byte, point geometry.Point) {
	if ok, next := isDeadEnd(maze, point); ok {
		delete(maze, point)
		fillDeadEnd(maze, next)
	}
}

func isDeadEnd(maze map[geometry.Point]byte, point geometry.Point) (bool, geometry.Point) {
	walls := 0
	var next geometry.Point
	for _, direction := range geometry.Directions {
		p := point.Move(direction)
		if _, ok := maze[p]; !ok {
			walls++
		} else {
			next = p
		}
	}
	return walls == 3 && maze[point] == OpenPassage, next
}

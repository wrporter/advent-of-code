package maze

import (
	"github.com/wrporter/advent-of-code-2019/internal/common/bytes"
	"github.com/wrporter/advent-of-code-2019/internal/common/geometry"
)

const (
	SolidWall   = '#'
	OpenPassage = '.'
)

const (
	Start = "AA"
	End   = "ZZ"
)

type DonutMaze struct {
	maze    map[geometry.Point]byte
	portals map[geometry.Point]geometry.Point
	start   geometry.Point
	end     geometry.Point
}

func New(maze []string) *DonutMaze {
	mazeMap := make(map[geometry.Point]byte)
	portals := make(map[string]geometry.Point)
	connections := make(map[geometry.Point]geometry.Point)

	for y, line := range maze {
		for x, char := range line {
			spot := byte(char)
			if spot == OpenPassage {
				mazeMap[geometry.NewPoint(x, y)] = spot
			}
			if isPortalPart(spot) {
				if y < len(maze)-1 && isPortalPart(maze[y+1][x]) {
					passage := geometry.NewPoint(x, y-1)
					if passage.Y-1 < 0 || maze[passage.Y-1][passage.X] != OpenPassage {
						passage = geometry.NewPoint(x, y+2)
					}

					name := string(maze[y][x]) + string(maze[y+1][x])
					if _, ok := portals[name]; ok {
						connections[portals[name]] = passage
						connections[passage] = portals[name]
					} else {
						portals[name] = passage
					}
				} else if x < len(maze[y])-1 && isPortalPart(maze[y][x+1]) {
					passage := geometry.NewPoint(x-1, y)
					if passage.X-1 < 0 || maze[passage.Y][passage.X-1] != OpenPassage {
						passage = geometry.NewPoint(x+2, y)
					}

					name := string(maze[y][x]) + string(maze[y][x+1])
					if _, ok := portals[name]; ok {
						connections[portals[name]] = passage
						connections[passage] = portals[name]
					} else {
						portals[name] = passage
					}
				}
			}
		}
	}

	return &DonutMaze{
		maze:    mazeMap,
		portals: connections,
		start:   portals[Start],
		end:     portals[End],
	}
}

type node struct {
	point geometry.Point
	steps int
}

func (m *DonutMaze) MinSteps() int {
	queue := []node{{m.start, 0}}
	visited := make(map[geometry.Point]bool)
	var cur node

	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		if cur.point == m.end {
			return cur.steps
		}
		for _, neighbor := range m.getNeighbors(cur.point, visited) {
			queue = append(queue, node{neighbor, cur.steps + 1})
		}
		visited[cur.point] = true
	}

	panic("Failed to find a route!")
}

func (m *DonutMaze) getNeighbors(point geometry.Point, visited map[geometry.Point]bool) []geometry.Point {
	var neighbors []geometry.Point
	if portal, ok := m.portals[point]; ok {
		neighbors = append(neighbors, portal)
	}
	for _, direction := range geometry.Directions {
		p := point.Add(direction)
		if !visited[p] {
			if m.maze[p] == OpenPassage {
				neighbors = append(neighbors, p)
			}
		}
	}
	return neighbors
}

func isPortalPart(b byte) bool {
	return bytes.IsLetter(b)
}

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
	portals map[geometry.Point]Portal
	start   geometry.Point
	end     geometry.Point
}

type Portal struct {
	name  string
	to    geometry.Point
	level int
}

const MaxDepth = 100

func New(maze []string) *DonutMaze {
	mazeMap := make(map[geometry.Point]byte)
	portals := make(map[string]geometry.Point)
	connections := make(map[geometry.Point]Portal)

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
						connections[portals[name]] = buildPortal(maze, name, portals[name], passage)
						connections[passage] = buildPortal(maze, name, passage, portals[name])
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
						connections[portals[name]] = buildPortal(maze, name, portals[name], passage)
						connections[passage] = buildPortal(maze, name, passage, portals[name])
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

func buildPortal(maze []string, name string, portal geometry.Point, to geometry.Point) Portal {
	level := 1
	if portal.Y == 2 ||
		portal.Y == len(maze)-3 ||
		portal.X == 2 ||
		portal.X == len(maze[portal.Y])-3 {
		level = -1
	}
	return Portal{name, to, level}
}

type node struct {
	point geometry.Point
	steps int
	level int
}

func (m *DonutMaze) MinSteps(recursive bool) int {
	queue := []node{{m.start, 0, 0}}
	visited := make(map[Neighbor]bool)
	var cur node

	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		visited[Neighbor{cur.point, cur.level}] = true

		if cur.point == m.end && cur.level == 0 {
			return cur.steps
		}

		for _, neighbor := range m.getNeighbors(cur, visited, recursive) {
			if neighbor.level >= 0 && neighbor.level < MaxDepth && !visited[neighbor] {
				queue = append(queue, node{neighbor.point, cur.steps + 1, neighbor.level})
			}
		}
	}

	return -1
}

type Neighbor struct {
	point geometry.Point
	level int
}

func (m *DonutMaze) getNeighbors(cur node, visited map[Neighbor]bool, recursive bool) []Neighbor {
	var neighbors []Neighbor
	if portal, ok := m.portals[cur.point]; ok {
		level := 0
		if recursive {
			level = cur.level + portal.level
		}
		neighbors = append(neighbors, Neighbor{portal.to, level})
	}
	for _, direction := range geometry.Directions {
		p := cur.point.Add(direction)
		if m.maze[p] == OpenPassage {
			neighbors = append(neighbors, Neighbor{p, cur.level})
		}
	}
	return neighbors
}

func isPortalPart(b byte) bool {
	return bytes.IsLetter(b)
}

package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/v2/dsu"
	"sort"
	"strings"
)

func part1(input string, args ...interface{}) interface{} {
	numConnections := args[0].(int)
	boxes, connections := parse(input)

	set := dsu.NewDSU(len(boxes))
	for i := 0; i < len(boxes) && i < numConnections; i++ {
		connection := connections[i]
		set.Union(connection.box1.id, connection.box2.id)
	}

	sizesMap := make(map[int]int)
	for i := range boxes {
		root := set.Find(i)
		sizesMap[root] = set.Size(root)
	}

	var sizes []int
	for _, s := range sizesMap {
		sizes = append(sizes, s)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
}

func part2(input string, _ ...interface{}) interface{} {
	boxes, connections := parse(input)

	set := dsu.NewDSU(len(boxes))
	numJunctions := len(boxes)
	var last1, last2 int

	for _, connection := range connections {
		isMerged := set.Union(connection.box1.id, connection.box2.id)
		if isMerged {
			numJunctions--
			if numJunctions == 1 {
				last1, last2 = connection.box1.id, connection.box2.id
				break
			}
		}
	}

	return boxes[last1].x * boxes[last2].x
}

func parse(input string) ([]*Box, []Connection) {
	lines := strings.Split(input, "\n")
	boxes := make([]*Box, len(lines))
	for i, line := range lines {
		values, _ := convert.ToInts(strings.Split(line, ","))
		boxes[i] = &Box{
			id: i,
			x:  values[0],
			y:  values[1],
			z:  values[2],
		}
	}

	var connections []Connection
	for i, box := range boxes {
		for j := i + 1; j < len(lines); j++ {
			connections = append(connections, Connection{
				box1:     box,
				box2:     boxes[j],
				distance: box.Distance(boxes[j]),
			})
		}
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})
	return boxes, connections
}

type Connection struct {
	box1, box2 *Box
	distance   int
}

type Box struct {
	id      int
	x, y, z int
}

func (b *Box) Distance(b2 *Box) int {
	dx := b.x - b2.x
	dy := b.y - b2.y
	dz := b.z - b2.z

	return dx*dx + dy*dy + dz*dz
}

func main() {
	New().Run([]interface{}{1000}, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2025, Day: 8, Part1: part1, Part2: part2}
}

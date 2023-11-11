package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/v2/geometry"
	"strconv"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	return moveRope(input, 2)
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	return moveRope(input, 10)
}

func moveRope(input string, size int) interface{} {
	rope := make([]*geometry.Point, size)
	for i := range rope {
		rope[i] = geometry.NewPoint(0, 0)
	}
	tail := rope[len(rope)-1]

	visited := map[geometry.Point]bool{*tail.Copy(): true}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		direction := directions[parts[0]]
		amount, _ := strconv.Atoi(parts[1])

		for step := 0; step < amount; step++ {
			rope[0].Move(direction)

			for knot := 1; knot < len(rope); knot++ {
				curr, prev := rope[knot], rope[knot-1]
				dx, dy := prev.Diff(curr)

				if curr.Distance(prev) > 1 {
					if dx < 0 {
						curr.X -= 1
					}
					if dx > 0 {
						curr.X += 1
					}
					if dy < 0 {
						curr.Y -= 1
					}
					if dy > 0 {
						curr.Y += 1
					}
				}
			}

			visited[*tail.Copy()] = true
		}

		//grid := geometry.MapToGrid(visited)
		//fmt.Println(line)
		//fmt.Println(rope[0], tail)
		//fmt.Println(geometry.RenderGrid(grid))
	}

	return len(visited)
}

var directions = map[string]geometry.Direction{
	"L": geometry.Left,
	"R": geometry.Right,
	"U": geometry.Up,
	"D": geometry.Down,
}

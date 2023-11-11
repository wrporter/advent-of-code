package solution

import (
	"fmt"
	mymath "github.com/wrporter/advent-of-code/internal/common/ints"
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
	rope := make([]*point, size)
	for i := range rope {
		rope[i] = &point{x: 0, y: 0}
	}
	tail := rope[len(rope)-1]

	visited := map[point]bool{tail.copy(): true}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		direction := directions[parts[0]]
		amount, _ := strconv.Atoi(parts[1])

		for step := 0; step < amount; step++ {
			rope[0].x += direction.x
			rope[0].y += direction.y

			for knot := 1; knot < len(rope); knot++ {
				curr, prev := rope[knot], rope[knot-1]
				d := prev.diff(curr)

				if curr.distance(prev) > 1 {
					if d.x < 0 {
						curr.x -= 1
					}
					if d.x > 0 {
						curr.x += 1
					}
					if d.y < 0 {
						curr.y -= 1
					}
					if d.y > 0 {
						curr.y += 1
					}
				}
			}

			visited[tail.copy()] = true
		}

		//grid := MapToGrid(visited)
		//fmt.Println(line)
		//fmt.Println(rope[0], tail)
		//fmt.Println(RenderGrid(grid))
	}

	return len(visited)
}

type point struct {
	x, y int
}

var directions = map[string]point{
	"L": {x: -1},
	"R": {x: 1},
	"U": {y: -1},
	"D": {y: 1},
}

func (p *point) diff(p2 *point) point {
	return point{
		x: p.x - p2.x,
		y: p.y - p2.y,
	}
}

func (p *point) manhattan(p2 *point) int {
	return mymath.Abs(p.x-p2.x) + mymath.Abs(p.y-p2.y)
}

func (p *point) distance(p2 *point) int {
	return mymath.Max(mymath.Abs(p.x-p2.x), mymath.Abs(p.y-p2.y))
}

func (p *point) copy() point {
	return point{x: p.x, y: p.y}
}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

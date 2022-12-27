package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/convert"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"math"
	"regexp"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 17
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var regex = regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)

func part1(input []string) interface{} {
	trench := parseInput(input[0])
	highest := math.MinInt
	tryVelocities(trench, func(h int) {
		highest = ints.Max(highest, h)
	})
	return highest
}

func part2(input []string) interface{} {
	trench := parseInput(input[0])
	count := 0
	tryVelocities(trench, func(_ int) {
		count++
	})
	return count
}

func tryVelocities(trench area, onReach func(highest int)) {
	for vx := 1; vx <= trench.br.x; vx++ {
		for vy := trench.br.y; vy <= -trench.br.y; vy++ {
			simulate(trench, vx, vy, onReach)
		}
	}
}

func simulate(trench area, vx int, vy int, onReach func(highest int)) {
	p := newProbe(0, 0, vx, vy)
	maxY := math.MinInt

	for t := 0; t <= 300; t++ {
		p.step()
		maxY = ints.Max(maxY, p.y)

		if trench.contains(*p.point) {
			onReach(maxY)
			break
		}
	}
}

func parseInput(line string) area {
	match := regex.FindStringSubmatch(line)
	x1, x2 := convert.StringToInt(match[1]), convert.StringToInt(match[2])
	y1, y2 := convert.StringToInt(match[3]), convert.StringToInt(match[4])
	return area{
		tl: point{x: ints.Min(x1, x2), y: ints.Max(y1, y2)},
		br: point{x: ints.Max(x1, x2), y: ints.Min(y1, y2)},
	}
}

func newProbe(x, y, vx, vy int) *probe {
	return &probe{
		point:           &point{x: x, y: y},
		velocity:        &point{x: vx, y: vy},
		initialVelocity: point{x: vx, y: vy},
	}
}

type probe struct {
	*point
	velocity        *point
	initialVelocity point
}

func (p *probe) step() {
	p.x += p.velocity.x
	p.y += p.velocity.y

	if p.velocity.x > 0 {
		p.velocity.x--
	} else if p.velocity.x < 0 {
		p.velocity.x++
	}

	p.velocity.y--
}

type point struct {
	x int
	y int
}

type area struct {
	tl point
	br point
}

func (a area) contains(p point) bool {
	return p.x >= a.tl.x && p.x <= a.br.x &&
		p.y <= a.tl.y && p.y >= a.br.y
}

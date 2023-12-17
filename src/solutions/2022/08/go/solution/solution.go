package solution

import (
	"aoc/src/lib/go/convert"
	mymath "aoc/src/lib/go/ints"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	trees := convert.ToIntGrid(input)
	visible := 2*len(trees) + 2*len(trees[0]) - 4

	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[y])-1; x++ {
			isVisible := false

			for i := 0; i < len(directions) && !isVisible; i++ {
				d := directions[i]
				ny := y + d.y
				nx := x + d.x

				for ny >= 0 && ny < len(trees) &&
					nx >= 0 && nx < len(trees[ny]) &&
					trees[y][x] > trees[ny][nx] &&
					!isVisible {

					if ny == 0 || ny == len(trees)-1 ||
						nx == 0 || nx == len(trees[ny])-1 {
						isVisible = true
					}

					ny += d.y
					nx += d.x
				}
			}

			if isVisible {
				visible += 1
			}
		}
	}

	return visible
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	trees := convert.ToIntGrid(input)
	maxScenicScore := 0

	for y := 1; y < len(trees)-1; y++ {
		for x := 1; x < len(trees[y])-1; x++ {
			viewingDistance := [4]int{1, 1, 1, 1}

			for i := 0; i < len(directions); i++ {
				done := false
				d := directions[i]
				ny := y + d.y
				nx := x + d.x

				for ny >= 0 && ny < len(trees) &&
					nx >= 0 && nx < len(trees[ny]) &&
					trees[y][x] > trees[ny][nx] &&
					!done {

					if ny == 0 || ny == len(trees)-1 ||
						nx == 0 || nx == len(trees[ny])-1 {
						done = true
					} else {
						viewingDistance[i] += 1
					}

					ny += d.y
					nx += d.x
				}
			}

			scenicScore := viewingDistance[0] * viewingDistance[1] * viewingDistance[2] * viewingDistance[3]
			maxScenicScore = mymath.Max(maxScenicScore, scenicScore)
		}
	}

	return maxScenicScore
}

type point struct {
	x, y int
}

var directions = [4]point{
	{x: -1},
	{x: 1},
	{y: -1},
	{y: 1},
}

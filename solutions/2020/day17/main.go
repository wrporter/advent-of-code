package main

import (
	"bytes"
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"time"
)

var debug = false

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 17
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	debug = false
	cube := parseInput(input)
	cube = run(cube, 6)
	return countAlive(cube)
}

func part2(input []string) interface{} {
	debug = false
	cube := parseInput2(input)
	cube = run2(cube, 6)
	return countAlive2(cube)
}

const (
	Alive = '#'
	Dead  = '.'
)

func run(cube [][][]rune, numCycles int) [][][]rune {
	printCycle(cube, 0)
	for i := 1; i <= numCycles; i++ {
		cube = cycle(cube)
		printCycle(cube, i)
	}
	return cube
}

func cycle(cube [][][]rune) [][][]rune {
	next := makeCube(len(cube) + 2)

	for z := range cube {
		for y := range cube[z] {
			for x, cell := range cube[z][y] {
				numAliveNeighbors := countAliveNeighbors(cube, z, y, x)
				if (cell == Alive && (numAliveNeighbors == 2 || numAliveNeighbors == 3)) ||
					(cell == Dead && numAliveNeighbors == 3) {
					next[z+1][y+1][x+1] = Alive
				}
			}
		}
	}

	return next
}

func countAlive(cube [][][]rune) int {
	count := 0

	for z := range cube {
		for y := range cube[z] {
			for x := range cube[z][y] {
				if cube[z][y][x] == Alive {
					count++
				}
			}
		}
	}

	return count
}

func countAliveNeighbors(cube [][][]rune, z, y, x int) int {
	count := 0

	for mz := -1; mz <= 1; mz++ {
		for my := -1; my <= 1; my++ {
			for mx := -1; mx <= 1; mx++ {
				// skip the current position
				if mz == 0 && my == 0 && mx == 0 {
					continue
				}

				nz := z + mz
				ny := y + my
				nx := x + mx

				// skip if out of bounds, assuming those are dead
				if nz < 0 || nz >= len(cube) ||
					ny < 0 || ny >= len(cube[nz]) ||
					nx < 0 || nx >= len(cube[nz][ny]) {
					continue
				}

				if cube[nz][ny][nx] == Alive {
					count++
				}
			}
		}
	}

	return count
}

func parseInput(input []string) [][][]rune {
	cube := makeCube(len(input) + 2)

	for y, row := range input {
		for x, cell := range row {
			cube[2][y+1][x+1] = cell
		}
	}

	return cube
}

func printCycle(cube [][][]rune, cycle int) {
	if !debug {
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Cycle = %d\n", cycle))

	for z := 1; z < len(cube)-1; z++ {
		buffer.WriteString(fmt.Sprintf("z = %d\n", z))
		for y := 1; y < len(cube[z])-1; y++ {
			for x := 1; x < len(cube[z][y])-1; x++ {
				buffer.WriteRune(cube[z][y][x])
			}
			buffer.WriteRune('\n')
		}
	}

	fmt.Println(buffer.String())
}

func makeCube(size int) [][][]rune {
	cube := make([][][]rune, size)

	for z := range cube {
		cube[z] = make([][]rune, size)
		for y := range cube[z] {
			cube[z][y] = make([]rune, size)
			for x := range cube[z][y] {
				cube[z][y][x] = Dead
			}
		}
	}

	return cube
}

func run2(cube [][][][]rune, numCycles int) [][][][]rune {
	printCycle2(cube, 0)
	for i := 1; i <= numCycles; i++ {
		cube = cycle2(cube)
		printCycle2(cube, i)
	}
	return cube
}

func cycle2(cube [][][][]rune) [][][][]rune {
	next := makeCube2(len(cube) + 2)

	for w := range cube {
		for z := range cube[w] {
			for y := range cube[w][z] {
				for x, cell := range cube[w][z][y] {
					numAliveNeighbors := countAliveNeighbors2(cube, w, z, y, x)
					if (cell == Alive && (numAliveNeighbors == 2 || numAliveNeighbors == 3)) ||
						(cell == Dead && numAliveNeighbors == 3) {
						next[w+1][z+1][y+1][x+1] = Alive
					}
				}
			}
		}
	}

	return next
}

func countAlive2(cube [][][][]rune) int {
	count := 0

	for w := range cube {
		for z := range cube[w] {
			for y := range cube[w][z] {
				for x := range cube[w][z][y] {
					if cube[w][z][y][x] == Alive {
						count++
					}
				}
			}
		}
	}

	return count
}

func countAliveNeighbors2(cube [][][][]rune, w, z, y, x int) int {
	count := 0

	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					// skip the current position
					if dw == 0 && dz == 0 && dy == 0 && dx == 0 {
						continue
					}

					nw := w + dw
					nz := z + dz
					ny := y + dy
					nx := x + dx

					// skip if out of bounds, assuming those are dead
					if nw < 0 || nw >= len(cube) ||
						nz < 0 || nz >= len(cube[nw]) ||
						ny < 0 || ny >= len(cube[nw][nz]) ||
						nx < 0 || nx >= len(cube[nw][nz][ny]) {
						continue
					}

					if cube[nw][nz][ny][nx] == Alive {
						count++
					}
				}
			}
		}
	}

	return count
}

func parseInput2(input []string) [][][][]rune {
	cube := makeCube2(len(input) + 2)

	for y, row := range input {
		for x, cell := range row {
			cube[2][2][y+1][x+1] = cell
		}
	}

	return cube
}

func printCycle2(cube [][][][]rune, cycle int) {
	if !debug {
		return
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Cycle = %d\n", cycle))

	for w := 1; w < len(cube)-1; w++ {
		for z := 1; z < len(cube[w])-1; z++ {
			buffer.WriteString(fmt.Sprintf("z = %d, w = %d\n", z, w))
			for y := 1; y < len(cube[w][z])-1; y++ {
				for x := 1; x < len(cube[w][z][y])-1; x++ {
					buffer.WriteRune(cube[w][z][y][x])
				}
				buffer.WriteRune('\n')
			}
		}
	}

	fmt.Println(buffer.String())
}

func makeCube2(size int) [][][][]rune {
	cube := make([][][][]rune, size)

	for w := range cube {
		cube[w] = make([][][]rune, size)
		for z := range cube[w] {
			cube[w][z] = make([][]rune, size)
			for y := range cube[w][z] {
				cube[w][z][y] = make([]rune, size)
				for x := range cube[w][z][y] {
					cube[w][z][y][x] = Dead
				}
			}
		}
	}

	return cube
}

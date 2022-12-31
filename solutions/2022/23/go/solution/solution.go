package solution

import (
	"github.com/wrporter/advent-of-code/internal/common/geometry"
	"strings"
)

func (s Solution) Part1(input string, _ ...interface{}) interface{} {
	elves := parseInput(input)
	result := moveElves(elves, 10)
	return result.sumEmptyTiles
}

func (s Solution) Part2(input string, _ ...interface{}) interface{} {
	elves := parseInput(input)
	result := moveElves(elves, 10_000)
	return result.round
}

func moveElves(grove map[geometry.Point]bool, maxRounds int) Result {
	firstDirection := 0
	anyElfHasMoved := false
	round := 1

	for ; round <= maxRounds; round++ {
		firstHalf := make(map[geometry.Point][]Move)

		for elf := range grove {
			isAdjacentElf := false
			for _, direction := range geometry.AllDirectionsModifiers {
				if grove[elf.Add(direction)] {
					isAdjacentElf = true
				}
			}

			hasMoved := false
			if isAdjacentElf {
				for group := firstDirection; group < firstDirection+len(directionGroups) && !hasMoved; group++ {
					directionGroup := directionGroups[group%len(directionGroups)]
					noElfInDirection := true
					for _, direction := range directionGroup {
						if grove[elf.Move(direction)] {
							noElfInDirection = false
						}
					}

					if noElfInDirection {
						hasMoved = true
						anyElfHasMoved = true

						direction := directionGroup[0]
						to := elf.Move(direction)
						entry := Move{
							from: elf,
							to:   to,
						}

						if _, exists := firstHalf[to]; exists {
							firstHalf[to] = append(firstHalf[to], entry)
						} else {
							firstHalf[to] = []Move{entry}
						}
					}
				}
			}

			if !hasMoved {
				firstHalf[elf] = []Move{{from: elf, to: elf}}
			}
		}

		secondHalf := make(map[geometry.Point]bool)
		for spot := range firstHalf {
			elves := firstHalf[spot]
			if len(elves) == 1 {
				secondHalf[elves[0].to] = true
			} else {
				for _, elf := range elves {
					secondHalf[elf.from] = true
				}
			}
		}

		grove = secondHalf
		firstDirection = (firstDirection + 1) % len(directionGroups)

		if !anyElfHasMoved {
			break
		}
		anyElfHasMoved = false
	}

	rectangle := geometry.MapToGrid(grove)
	sumEmptyTiles := 0
	for y := 0; y < len(rectangle); y++ {
		for x := 0; x < len(rectangle[y]); x++ {
			if rectangle[y][x] == '.' {
				sumEmptyTiles += 1
			}
		}
	}

	//fmt.Println(geometry.RenderGrid(rectangle))

	return Result{sumEmptyTiles: sumEmptyTiles, round: round}
}

type Result struct {
	sumEmptyTiles int
	round         int
}

var directionGroups = [][]geometry.Direction{
	{geometry.Up, geometry.UpRight, geometry.UpLeft},
	{geometry.Down, geometry.DownRight, geometry.DownLeft},
	{geometry.Left, geometry.UpLeft, geometry.DownLeft},
	{geometry.Right, geometry.UpRight, geometry.DownRight},
}

func parseInput(input string) map[geometry.Point]bool {
	elves := make(map[geometry.Point]bool)
	for y, line := range strings.Split(input, "\n") {
		for x := 0; x < len(line); x++ {
			if string(line[x]) == "#" {
				elves[geometry.Point{
					X: x,
					Y: y,
				}] = true
			}
		}
	}
	return elves
}

type Move struct {
	from geometry.Point
	to   geometry.Point
}

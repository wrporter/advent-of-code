package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/stringgrid"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 20
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	tiles := parse(input)
	image := getCorrectOrientation(tiles)

	cornerTiles := []int{
		image[0][0].ID,
		image[0][len(image[0])-1].ID,
		image[len(image)-1][0].ID,
		image[len(image)-1][len(image[0])-1].ID,
	}
	return ints.Product(cornerTiles)
}

func part2(input []string) interface{} {
	tiles := parse(input)
	tileImage := getCorrectOrientation(tiles)
	image := assembleImage(tileImage)

	monster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	numMonsters := countMonsters(image, monster)

	numPoundSignsInImage := strings.Count(strings.Join(image, ""), "#")
	numPoundSignsInMonster := strings.Count(strings.Join(monster, ""), "#")
	numNotMonsters := numPoundSignsInImage - (numPoundSignsInMonster * numMonsters)
	return numNotMonsters
}

func countMonsters(image []string, monster []string) int {
	monsterPatterns := getArrangements(monster)

	for _, pattern := range monsterPatterns {
		count := 0
		for dr := 0; dr <= len(image)-len(pattern); dr++ {
			for dc := 0; dc <= len(image[0])-len(pattern[0]); dc++ {
				if matches(image, pattern, dr, dc) {
					count++
				}
			}
		}

		if count != 0 {
			return count
		}
	}

	return 0
}

func matches(image []string, pattern []string, dr int, dc int) bool {
	for r := range pattern {
		for c := range pattern[r] {
			if pattern[r][c] == '#' && image[r+dr][c+dc] != '#' {
				return false
			}
		}
	}
	return true
}

func assembleImage(image [][]*Tile) []string {
	tileSize := len(image[0][0].Grid) - 2
	size := len(image) * tileSize
	result := make([]string, size)

	for imageRow := range image {
		for _, tile := range image[imageRow] {
			for row := 1; row < len(tile.Grid)-1; row++ {
				rowDelta := (imageRow * tileSize) + (row - 1)
				result[rowDelta] += tile.Grid[row][1 : len(tile.Grid)-1]
			}
		}
	}

	return result
}

func getCorrectOrientation(tiles map[int]*Tile) [][]*Tile {
	imageSize := ints.Sqrt(len(tiles))
	image := make([][]*Tile, imageSize)
	for i := range image {
		image[i] = make([]*Tile, imageSize)
	}
	used := make(map[int]bool)
	return findCorrectOrientation(image, 0, 0, tiles, used)
}

func findCorrectOrientation(image [][]*Tile, row int, col int, tiles map[int]*Tile, used map[int]bool) [][]*Tile {
	if row == len(image) {
		return image
	}

	if col == len(image[row]) {
		return findCorrectOrientation(image, row+1, 0, tiles, used)
	}

	for tileID, tile := range tiles {
		if used[tileID] {
			continue
		}

		for _, candidate := range tile.Arrangements {
			if row == 0 && col == 0 ||
				(row > 0 && stringgrid.Top(candidate) == stringgrid.Bottom(image[row-1][col].Grid)) ||
				(col > 0 && stringgrid.Left(candidate) == stringgrid.Right(image[row][col-1].Grid)) {

				image[row][col] = NewTile(tileID, candidate)
				used[tileID] = true

				result := findCorrectOrientation(image, row, col+1, tiles, used)
				if isValid(result) {
					return result
				}

				used[tileID] = false
				image[row][col] = nil
			}
		}
	}
	return image
}

func isValid(image [][]*Tile) bool {
	for _, row := range image {
		for _, tile := range row {
			if tile == nil {
				return false
			}
		}
	}
	return true
}

type (
	Tile struct {
		ID           int
		Grid         []string
		Arrangements [][]string
	}
)

func NewTile(id int, grid []string) *Tile {
	return &Tile{
		ID:   id,
		Grid: grid,
	}
}

func addArrangements(tile *Tile) {
	tile.Arrangements = getArrangements(tile.Grid)
}

func getArrangements(grid []string) [][]string {
	var a [][]string
	a = append(a, grid)
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.FlipLR(grid))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	a = append(a, stringgrid.Rotate90Clockwise(a[len(a)-1]))
	return a
}

var tileIDRegex = regexp.MustCompile(`^Tile (\d+):$`)

func parse(input []string) map[int]*Tile {
	tiles := make(map[int]*Tile)

	var tileID int
	var grid []string
	for i, line := range input {
		if line == "" || i == len(input)-1 {
			if i == len(input)-1 {
				grid = append(grid, line)
			}
			tile := NewTile(tileID, grid)
			addArrangements(tile)
			tiles[tileID] = tile
			grid = nil
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			match := tileIDRegex.FindStringSubmatch(line)
			tileID = conversion.StringToInt(match[1])
		} else {
			grid = append(grid, line)
		}
	}

	return tiles
}

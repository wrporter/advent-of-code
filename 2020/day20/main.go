package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/mystrings"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/stringgrid"
	"regexp"
	"strings"
)

func main() {
	year, day := 2020, 20
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var tileIDRegex = regexp.MustCompile(`^Tile (\d+):$`)

func part1(input []string) interface{} {
	tiles := parse(input)

	for _, aTile := range tiles {
		for _, bTile := range tiles {
			if aTile.ID != bTile.ID {
				for _, border := range bTile.Borders {
					if stringgrid.Contains(aTile.Borders, border) ||
						stringgrid.Contains(aTile.FlippedBorders, border) {
						aTile.BorderTiles = append(aTile.BorderTiles, bTile.ID)
					}
				}
			}
		}
	}

	var cornerTiles []int
	for _, tile := range tiles {
		if len(tile.BorderTiles) == 2 {
			cornerTiles = append(cornerTiles, tile.ID)
		}
	}

	// [1831, 1699, 2309, 2789]
	return ints.Product(cornerTiles)
}

func part2(input []string) interface{} {
	return 0
}

type (
	Tile struct {
		ID             int
		Grid           []string
		Borders        []string
		FlippedBorders []string
		BorderTiles    []int
	}
)

func NewTile(id int, grid []string) *Tile {
	tile := &Tile{
		ID:   id,
		Grid: grid,
	}

	tile.Borders = []string{
		grid[0],
		grid[len(grid)-1],
		stringgrid.GetCol(grid, 0),
		stringgrid.GetCol(grid, len(grid)-1),
	}

	for _, border := range tile.Borders {
		tile.FlippedBorders = append(tile.FlippedBorders, mystrings.Reverse(border))
	}

	return tile
}

func parse(input []string) map[int]*Tile {
	tiles := make(map[int]*Tile)

	var tileID int
	var grid []string
	for i, line := range input {
		if line == "" || i == len(input)-1 {
			if i == len(input)-1 {
				grid = append(grid, line)
			}
			tiles[tileID] = NewTile(tileID, grid)
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

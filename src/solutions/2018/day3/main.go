package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/geometry"
	"aoc/src/lib/go/out"
	"fmt"
	"regexp"
)

func main() {
	year, day := 2018, 3
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	claims := parseClaims(input)

	overlaps := make(map[geometry.Point]int)
	for _, claim := range claims {
		for x := claim.StartX; x < (claim.StartX + claim.Width); x++ {
			for y := claim.StartY; y < (claim.StartY + claim.Height); y++ {
				point := geometry.NewPoint(x, y)
				overlaps[point]++
			}
		}
	}

	overlappingSquareInches := 0
	for _, count := range overlaps {
		if count >= 2 {
			overlappingSquareInches++
		}
	}

	return overlappingSquareInches
}

func part2(input []string) interface{} {
	claims := parseClaims(input)

	overlappingClaims := make(map[int][]int)
	claimsOnPoint := make(map[geometry.Point][]int)
	for claimID, claim := range claims {
		for x := claim.StartX; x < (claim.StartX + claim.Width); x++ {
			for y := claim.StartY; y < (claim.StartY + claim.Height); y++ {
				point := geometry.NewPoint(x, y)
				if claimIDs, ok := claimsOnPoint[point]; ok {
					for _, otherClaimID := range claimIDs {
						overlappingClaims[claimID] = append(overlappingClaims[claimID], otherClaimID)
						overlappingClaims[otherClaimID] = append(overlappingClaims[otherClaimID], claimID)
					}
				}
				claimsOnPoint[point] = append(claimsOnPoint[point], claimID)
			}
		}
	}

	for claimID := range claims {
		if _, ok := overlappingClaims[claimID]; !ok {
			return claimID
		}
	}

	return 0
}

var regex = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

type Claim struct {
	StartX int
	StartY int
	Width  int
	Height int
}

func parseClaims(input []string) map[int]Claim {
	claims := make(map[int]Claim)
	for _, line := range input {
		match := regex.FindStringSubmatch(line)
		claims[convert.StringToInt(match[1])] = Claim{
			StartX: convert.StringToInt(match[2]),
			StartY: convert.StringToInt(match[3]),
			Width:  convert.StringToInt(match[4]),
			Height: convert.StringToInt(match[5]),
		}
	}
	return claims
}

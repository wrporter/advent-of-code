package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"strings"
	"unicode"
)

func main() {
	year, day := 2018, 5
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

//     buf = []
//    for c in line:
//        if buf and are_opp(c, buf[-1]):
//            buf.pop()
//        else:
//            buf.append(c)
//    return len(buf)

func part1(input []string) interface{} {
	return len(react(input[0]))
}

func part2(input []string) interface{} {
	polymer := input[0]

	unitTypes := make(map[rune]bool)
	for _, char := range polymer {
		unitTypes[unicode.ToLower(char)] = true
	}

	min := ints.MaxInt
	for unitType := range unitTypes {
		result := strings.ReplaceAll(polymer, string(unitType), "")
		result = strings.ReplaceAll(result, string(unicode.ToUpper(unitType)), "")
		result = react(result)
		min = ints.Min(min, len(result))
	}

	return min
}

func react(line string) string {
	var polymer []rune

	for _, char := range line {
		if len(polymer) > 0 && oppositePolarity(char, polymer[len(polymer)-1]) {
			polymer = polymer[:len(polymer)-1]
		} else {
			polymer = append(polymer, char)
		}
	}

	return string(polymer)
}

func oppositePolarity(char1 rune, char2 rune) bool {
	return unicode.ToUpper(char1) == unicode.ToUpper(char2) && (unicode.IsUpper(char1) && unicode.IsLower(char2) ||
		unicode.IsLower(char1) && unicode.IsUpper(char2))
}

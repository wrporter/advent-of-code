package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	year, day := 2020, 14
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

var (
	maskRegex = regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRegex  = regexp.MustCompile(`^mem\[(\d+)] = (\d+)$`)
)

func part1(input []string) interface{} {
	memory := make(map[int]int)
	var mask string
	for _, line := range input {
		if strings.Contains(line, "mask") {
			match := maskRegex.FindStringSubmatch(line)
			mask = match[1]
		} else if strings.Contains(line, "mem") {
			match := memRegex.FindStringSubmatch(line)
			address := conversion.StringToInt(match[1])
			value := conversion.StringToInt(match[2])

			for i, bit := range mask {
				position := len(mask) - i - 1
				if bit == '1' {
					value = setBit(value, position)
				} else if bit == '0' {
					value = clearBit(value, position)
				}
			}

			memory[address] = value
		}
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func part2(input []string) interface{} {
	memory := make(map[int]int)
	var mask string
	for _, line := range input {
		if strings.Contains(line, "mask") {
			match := maskRegex.FindStringSubmatch(line)
			mask = match[1]
		} else if strings.Contains(line, "mem") {
			match := memRegex.FindStringSubmatch(line)
			address := conversion.StringToInt(match[1])
			value := conversion.StringToInt(match[2])

			numXPermutations := getNumXPermutations(mask)
			newAddress := address

			for x := 0; x < numXPermutations; x++ {
				newMask := getMask(mask, x)

				for i, bit := range mask {
					position := len(mask) - i - 1
					if bit == '1' {
						newAddress = setBit(newAddress, position)
					} else if bit == 'X' {
						maskBit := newMask[i]
						if maskBit == '1' {
							newAddress = setBit(newAddress, position)
						} else if maskBit == '0' {
							newAddress = clearBit(newAddress, position)
						}
					}
				}
				memory[newAddress] = value
			}
		}
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func getNumXPermutations(mask string) int {
	numXs := strings.Count(mask, "X")
	value := 1
	for i := 0; i < numXs; i++ {
		value += ints.Pow(2, i)
	}
	return value
}

func getMask(mask string, x int) string {
	result := make([]rune, len(mask))
	numXs := strings.Count(mask, "X")
	xBits := toBits(x, numXs)
	xi := len(xBits) - 1
	for i := 0; i < len(mask); i++ {
		if mask[i] == 'X' {
			result[i] = rune(xBits[xi])
			xi--
		} else {
			result[i] = rune(mask[i])
		}
	}
	return string(result)
}

func toBits(value int, size int) string {
	bits := strconv.FormatInt(int64(value), 2)
	result := bits
	for i := 0; i < size-len(bits); i++ {
		result = "0" + result
	}
	return result
}

func setBit(n int, pos int) int {
	n |= 1 << pos
	return n
}

func clearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/conversion"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"github.com/wrporter/advent-of-code/internal/common/ints"
	"github.com/wrporter/advent-of-code/internal/common/out"
	"github.com/wrporter/advent-of-code/internal/common/timeit"
	"strings"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 10
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	lengths, _ := conversion.ToInts(strings.Split(input[0], ","))
	return singleRoundHash(256, lengths)
}

func singleRoundHash(size int, lengths []int) int {
	var list []int
	for value := 0; value < size; value++ {
		list = append(list, value)
	}

	position := 0
	skip := 0

	for _, length := range lengths {
		list = reverse(list, position, length)

		position = (position + length + skip) % size
		skip++
	}

	return list[0] * list[1]
}

func reverse(list []int, position int, length int) []int {
	result := ints.Copy(list)
	size := len(result)

	i := position
	j := (position + length - 1) % size

	for iterations := 0; iterations < length/2; iterations++ {
		result[i], result[j] = result[j], result[i]
		i, j = ints.WrapMod(i+1, size), ints.WrapMod(j-1, size)
	}
	return result
}

func part2(input []string) interface{} {
	lengths := []rune(input[0])
	suffix := []rune{17, 31, 73, 47, 23}
	lengths = append(lengths, suffix...)
	result := hash(256, lengths)

	return result
}

func hash(size int, lengths []rune) string {
	var list []int
	for value := 0; value < size; value++ {
		list = append(list, value)
	}

	position := 0
	skip := 0

	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			list = reverse(list, position, int(length))

			position = (position + int(length) + skip) % size
			skip++
		}
	}

	dense := denseHash(list)
	hex := toHexString(dense)

	return hex
}

func toHexString(values []int) string {
	hex := ""
	for _, value := range values {
		chunk := fmt.Sprintf("%x", value)
		if len(chunk) == 1 {
			chunk = "0" + chunk
		}
		hex += chunk
	}
	return hex
}

func denseHash(values []int) []int {
	var result []int
	for i := 0; i < len(values); i += 16 {
		block := values[i]
		for j := 1; j < 16; j++ {
			block ^= values[i+j]
		}
		result = append(result, block)
	}
	return result
}

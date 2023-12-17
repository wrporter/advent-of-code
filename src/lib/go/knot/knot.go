package knot

import (
	"aoc/src/lib/go/ints"
	"fmt"
)

func Hash(str string) string {
	lengths := []rune(str)
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
			list = Reverse(list, position, int(length))

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

func Reverse(list []int, position int, length int) []int {
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

package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2020, 25
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	publicKeys, _ := convert.ToInts(input)
	cardPublicKey := publicKeys[0]
	doorPublicKey := publicKeys[1]

	subjectNumber := 7
	cardLoopSize := findLoopSize(cardPublicKey, subjectNumber)
	doorLoopSize := findLoopSize(doorPublicKey, subjectNumber)

	cardEncryptionKey := getEncryptionKey(cardPublicKey, doorLoopSize)
	doorEncryptionKey := getEncryptionKey(doorPublicKey, cardLoopSize)

	if cardEncryptionKey == doorEncryptionKey {
		return cardEncryptionKey
	}

	return 0
}

func findLoopSize(key int, subjectNumber int) int {
	value := 1
	for loopSize := 1; loopSize < 1_000_000_000; loopSize++ {
		value *= subjectNumber
		value %= 20201227

		if value == key {
			return loopSize
		}
	}
	return 0
}

func getEncryptionKey(subjectNumber int, loopSize int) int {
	value := 1
	for loop := 1; loop <= loopSize; loop++ {
		value *= subjectNumber
		value %= 20201227
	}
	return value
}

func part2(input []string) interface{} {
	return "Merry Christmas!!!"
}

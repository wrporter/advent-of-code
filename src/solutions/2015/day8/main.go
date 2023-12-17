package main

import (
	"aoc/src/lib/go/file"
	"fmt"
	"strconv"
)

func countMemoryCharDiff(stringLiterals []string) int {
	numCodeChars := 0
	numMemoryChars := 0

	for _, stringLiteral := range stringLiterals {
		numCodeChars += len(stringLiteral)
		memoryString, _ := strconv.Unquote(stringLiteral)
		numMemoryChars += len(memoryString)
	}

	return numCodeChars - numMemoryChars
}

func countEncodedCharDiff(stringLiterals []string) int {
	numCodeChars := 0
	numEncodedChars := 0

	for _, stringLiteral := range stringLiterals {
		numCodeChars += len(stringLiteral)
		encodedString := strconv.Quote(stringLiteral)
		numEncodedChars += len(encodedString)
	}

	return numEncodedChars - numCodeChars
}

func main() {
	lines, _ := file.ReadFile("./2015/day8/input.txt")
	fmt.Println(countMemoryCharDiff(lines))
	fmt.Println(countEncodedCharDiff(lines))
}

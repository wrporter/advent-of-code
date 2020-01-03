package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/file"
	"strconv"
)

func countCharacterDiff(stringLiterals []string) int {
	numCodeChars := 0
	numMemoryChars := 0

	for _, stringLiteral := range stringLiterals {
		numCodeChars += len(stringLiteral)
		memoryString, _ := strconv.Unquote(stringLiteral)
		numMemoryChars += len(memoryString)
	}

	return numCodeChars - numMemoryChars
}

func main() {
	lines, _ := file.ReadFile("./2015/day8/input.txt")
	fmt.Println(countCharacterDiff(lines))
}

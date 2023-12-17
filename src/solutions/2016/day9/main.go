package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"fmt"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`\((\d+)x(\d+)\)`)

func part1(file string) int {
	for i := 0; i < len(file); i++ {
		if file[i] == '(' {
			marker := ""
			lastIndex := i
			for j := i; j == 0 || file[j-1] != ')'; j++ {
				marker += string(file[j])
				lastIndex = j
			}
			match := regex.FindStringSubmatch(marker)
			numChars := convert.StringToInt(match[1])
			repeat := convert.StringToInt(match[2])

			repeatChars := strings.Repeat(file[lastIndex+1:lastIndex+1+numChars], repeat)
			file = file[:i] + repeatChars + file[lastIndex+1+numChars:]
			i = i - 1 + len(repeatChars)
		}
	}

	return len(file)
}

func part2(file string) int {
	if !strings.ContainsRune(file, '(') {
		return len(file)
	}
	length := 0

	for strings.ContainsRune(file, '(') {
		startMarkerIndex := strings.IndexRune(file, '(')
		endMarkerIndex := strings.IndexRune(file, ')')

		marker := strings.Split(file[startMarkerIndex+1:endMarkerIndex], "x")
		numChars := convert.StringToInt(marker[0])
		repeat := convert.StringToInt(marker[1])

		length += startMarkerIndex + (part2(file[endMarkerIndex+1:endMarkerIndex+1+numChars]) * repeat)

		file = file[endMarkerIndex+1+numChars:]
	}

	length += len(file)
	return length
}

func main() {
	input, _ := file.ReadFile("./2016/day9/input.txt")
	answer1 := part1(input[0])
	answer2 := part2(input[0])
	fmt.Println(answer1)
	fmt.Println(answer2)
}

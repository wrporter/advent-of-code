package main

import (
	"aoc/src/lib/go/convert"
	"aoc/src/lib/go/file"
	"aoc/src/solutions/2019/day5/public/computer"
	"fmt"
	"strings"
)

func main() {
	codeLines, _ := file.ReadFile("./2019/day9/input.txt")
	memory, _ := convert.ToInts(strings.Split(codeLines[0], ","))
	//memory := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	//memory := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	//memory := []int{104, 1125899906842624, 99}
	cpu := computer.New()
	program := computer.NewProgram(memory)
	program.Input <- 2
	go cpu.RunProgram(program)

	var actualOutput []int
	for value := range program.Output {
		actualOutput = append(actualOutput, value)
	}
	fmt.Println(actualOutput)
}

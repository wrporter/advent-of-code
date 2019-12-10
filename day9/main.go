package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day5/public/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	codeLines, _ := file.ReadFile("./day9/input.txt")
	memory, _ := conversion.ToInts(strings.Split(codeLines[0], ","))
	//memory := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	//memory := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	//memory := []int{104, 1125899906842624, 99}
	cpu := computer.New()
	program := computer.NewProgram(memory)
	program.Input <- 1
	go cpu.RunProgram(program)

	var actualOutput []int
	for value := range program.Output {
		actualOutput = append(actualOutput, value)
	}
	fmt.Println(actualOutput)
}

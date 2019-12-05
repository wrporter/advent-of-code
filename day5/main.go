package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/day5/internal/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"github.com/wrporter/advent-of-code-2019/internal/common/file"
	"strings"
)

func main() {
	codeLines, _ := file.ReadFile("./day5/input.txt")
	//codeLines := []string{"3,0,4,0,99"}
	//codeLines := []string{"1002,4,3,4,33"}
	//codeLines := []string{"1101,100,-1,4,0"}
	//codeLines := []string{"3,9,8,9,10,9,4,9,99,-1,8"}
	//codeLines := []string{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"}
	program, _ := conversion.ToInts(strings.Split(codeLines[0], ","))
	cpu := computer.New()
	fmt.Println(cpu.Run(program, 1))
}

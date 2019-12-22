package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/intcode"
	"math"
)

//.................
//.................
//@................
//#####..#.########

type SpringDroid struct {
	code []int
}

var surveyHullCommands = []string{
	"NOT A J",
	"NOT C T",
	"AND D T",
	"OR T J",
	"WALK",
}

func (d *SpringDroid) SurveyHull() {
	cpu := intcode.New()
	program := intcode.NewProgram(d.code)
	cpu.Run(program)
	fmt.Println(receive(program.Output))

	for _, command := range surveyHullCommands {
		send(program.Input, command)
	}

	for output := range program.Output {
		if output > math.MaxInt8 {
			fmt.Printf("%d\n", output)
		} else {
			fmt.Printf("%c", output)
		}
	}
}

func receive(ch <-chan int) string {
	var data []byte
	for char := byte(<-ch); char != '\n'; char = byte(<-ch) {
		data = append(data, char)
	}
	return string(data)
}

func send(ch chan<- int, command string) {
	fmt.Println(command)
	for _, char := range command {
		ch <- int(char)
	}
	ch <- '\n'
}

func main() {
	code := intcode.ReadCode("./day21/input.txt")
	droid := &SpringDroid{code}
	droid.SurveyHull()
}

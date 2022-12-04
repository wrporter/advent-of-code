package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/intcode"
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
	"NOT A J", // if no hull 1 tile away, jump
	"NOT C T",
	"AND D T",
	"OR T J", // if no hull 3 or 4 tiles away, jump
	"WALK",
}

var surveyEntireHullCommands = []string{
	"NOT C T",
	"NOT B J",
	"OR T J",
	"NOT A T",
	"OR T J", // if no hull 1 or 2 or 3 tiles away, jump
	"OR E T",
	"OR H T",
	"AND D T",
	"AND T J", // if hull 4 tiles away, jump
	"RUN",
}

func (d *SpringDroid) SurveyHull(commands []string) {
	cpu := intcode.New()
	program := intcode.NewProgram(d.code)
	cpu.Run(program)
	fmt.Println(receive(program.Output))

	for _, command := range commands {
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
	droid.SurveyHull(surveyHullCommands)
	droid.SurveyHull(surveyEntireHullCommands)
}

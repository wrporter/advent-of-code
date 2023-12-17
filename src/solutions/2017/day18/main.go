package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2017, 18
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	registers := make(map[string]int)
	var sound int

	for i := 0; i < len(input); {
		command := strings.Fields(input[i])
		switch command[0] {
		case "snd":
			sound = getValue(registers, command[1])
		case "set":
			registers[command[1]] = getValue(registers, command[2])
		case "add":
			registers[command[1]] += getValue(registers, command[2])
		case "mul":
			registers[command[1]] *= getValue(registers, command[2])
		case "mod":
			registers[command[1]] %= getValue(registers, command[2])
		case "rcv":
			if getValue(registers, command[1]) != 0 {
				if sound != 0 {
					return sound
				}
			}
		case "jgz":
			if getValue(registers, command[1]) > 0 {
				i += getValue(registers, command[2])
				continue
			}
		}
		i++
	}

	return "fail"
}

func part2(input []string) interface{} {
	program0 := NewProgram(0, input)
	program1 := NewProgram(1, input)
	program0.Input = program1.Output
	program1.Input = program0.Output

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go program0.Run(wg)
	go program1.Run(wg)
	wg.Wait()

	return program1.sendCount
}

type Program struct {
	ID        int
	Registers map[string]int
	Input     chan int
	Output    chan int
	position  int
	commands  []string
	sendCount int
}

func NewProgram(id int, commands []string) *Program {
	output := make(chan int, 10000)
	registers := map[string]int{"p": id}

	return &Program{
		ID:        id,
		Registers: registers,
		Output:    output,
		position:  0,
		commands:  commands,
		sendCount: 0,
	}
}

func (p *Program) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(p.Output)
	//defer func() {
	//	fmt.Printf("[%d] Stopped\n", p.ID)
	//}()

	for p.position >= 0 && p.position < len(p.commands) {
		command := strings.Fields(p.commands[p.position])
		args := command[1:]

		switch command[0] {
		case "snd":
			//fmt.Printf("[%d] Send: %s\n", p.ID, command)
			p.Output <- getValue(p.Registers, args[0])
			p.sendCount++
		case "set":
			p.Registers[args[0]] = getValue(p.Registers, args[1])
		case "add":
			p.Registers[args[0]] += getValue(p.Registers, args[1])
		case "mul":
			p.Registers[args[0]] *= getValue(p.Registers, args[1])
		case "mod":
			p.Registers[args[0]] %= getValue(p.Registers, args[1])
		case "rcv":
			select {
			case <-time.After(250 * time.Millisecond):
				//fmt.Printf("[%d] Timeout\n", p.ID)
				return
			case p.Registers[args[0]] = <-p.Input:
				//fmt.Printf("[%d] Receive: %s\n", p.ID, command)
			}
		case "jgz":
			if getValue(p.Registers, args[0]) > 0 {
				p.position += getValue(p.Registers, args[1])
				continue
			}
		}
		p.position++
	}
}

func getValue(registers map[string]int, registerOrInt string) int {
	if value, err := strconv.Atoi(registerOrInt); err == nil {
		return value
	}
	return registers[registerOrInt]
}

package main

import (
	"bufio"
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/intcode"
	"github.com/wrporter/advent-of-code-2019/internal/common/probability"
	"os"
	"strings"
)

type Droid struct {
	program *intcode.Program
}

func New(code []int) *Droid {
	return &Droid{intcode.NewProgram(code)}
}

var exploreCommands = []string{
	"west",
	"take ornament",
	"south",
	"east",
	"take weather machine",
	"west",
	"north",
	"west",
	"take astrolabe",
	"north",
	"take fuel cell",
	"south",
	"south",
	"take hologram",
	"north",
	"east",
	"east",
	"east",
	"take mug",
	"north",
	"take monolith",
	"south",
	"south",
	"west",
	"north",
	"west",
	"take bowl of rice",
	"north",
	"west",
}

var requiredItems = []string{
	"fuel cell",
	"astrolabe",
	"ornament",
	"hologram",
}

var safeItems = []string{
	"ornament",
	"weather machine",
	"astrolabe",
	"fuel cell",
	"hologram",
	"mug",
	"monolith",
	"bowl of rice",
}

func (d *Droid) Deploy(manual bool) {
	cpu := intcode.New()
	cpu.Run(d.program)
	reader := bufio.NewReader(os.Stdin)
	var inventory []string
	explorePointer := 0

	for {
		for {
			message := receive(d.program.Output)
			fmt.Println(message)
			if message == "Command?" {
				break
			}
		}
		if manual {
			command, _ := reader.ReadString('\n')
			send(d.program.Input, command[:len(command)-1])
		} else {
			if explorePointer < len(exploreCommands) {
				command := exploreCommands[explorePointer]
				if strings.Contains(command, "take") {
					inventory = append(inventory, command[strings.LastIndex(command, "take ")+5:])
				}
				send(d.program.Input, command)
				explorePointer++
			} else {
				d.drop(inventory)

				// Try all combinations of items until the droid is the correct weight
				probability.ComboSpots(safeItems, 1, len(safeItems), func(candidateItems []string) {
					d.take(candidateItems)
					sendAndReceive(d.program, "inv")
					sendAndReceive(d.program, "north")
					d.drop(candidateItems)
				})
			}
		}
	}
}

func (d *Droid) take(items []string) {
	for _, item := range items {
		sendAndReceive(d.program, "take "+item)
	}
}

func (d *Droid) drop(items []string) {
	for _, item := range items {
		sendAndReceive(d.program, "drop "+item)
	}
}

func sendAndReceive(program *intcode.Program, command string) {
	send(program.Input, command)
	for {
		message := receive(program.Output)
		fmt.Println(message)
		if message == "Command?" {
			break
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
	for _, char := range command {
		ch <- int(char)
	}
	ch <- '\n'
}

func main() {
	code := intcode.ReadCode("./day25/input.txt")
	droid := New(code)
	droid.Deploy(true)
}

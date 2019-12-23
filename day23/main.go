package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code-2019/internal/common/intcode"
)

type Packet struct {
	x int
	y int
}

type Network struct {
	core    *intcode.Computer
	code    []int
	network map[int]*intcode.Program
}

type Computer struct {
	program *intcode.Program
}

func New(code []int) *Network {
	return &Network{
		core:    intcode.New(),
		code:    code,
		network: make(map[int]*intcode.Program),
	}
}

func (n *Network) BootComputers() {
	for address := 0; address < 50; address++ {
		computer := intcode.NewProgram(n.code)
		n.network[address] = computer
		n.core.Run(computer)
		computer.Input <- address
	}

	for i := 0; ; i = (i + 1) % 50 {
		computer := n.network[i]
		select {
		case address := <-computer.Output:
			x := <-computer.Output
			y := <-computer.Output
			if address == 255 {
				fmt.Printf("255th packet: (x: %d, y: %d)", x, y)
				return
			}
			n.network[address].Input <- x
			n.network[address].Input <- y
		case computer.Input <- -1:
		}
	}
}

func main() {
	code := intcode.ReadCode("./day23/input.txt")
	network := New(code)
	network.BootComputers()
}

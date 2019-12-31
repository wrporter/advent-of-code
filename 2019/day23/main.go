package main

import (
	"fmt"
	"github.com/wrporter/advent-of-code/internal/common/intcode"
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

func New(code []int) *Network {
	return &Network{
		core:    intcode.New(),
		code:    code,
		network: make(map[int]*intcode.Program),
	}
}

func (n *Network) Boot() {
	for address := 0; address < 50; address++ {
		computer := intcode.NewProgram(n.code)
		n.network[address] = computer
		n.core.Run(computer)
		computer.Input <- address
		computer.Input <- -1
	}

	var prevNat Packet
	var nat Packet
	idleComputers := 0
	numNatReceived := 0

	for i := 0; ; i = (i + 1) % 50 {
		computer := n.network[i]
		select {
		case address := <-computer.Output:
			x := <-computer.Output
			y := <-computer.Output
			//fmt.Printf("Packet Sent (from: %d, address: %d, x: %d, y: %d)\n", i, address, x, y)

			if address == 255 {
				numNatReceived++
				if numNatReceived == 1 {
					fmt.Printf("First packet sent to address 255: (x: %d, y: %d)\n", x, y)
				}
				nat = Packet{x, y}
			} else {
				n.network[address].Input <- x
				n.network[address].Input <- y
			}

			idleComputers = 0
		case computer.Input <- -1:
			//fmt.Printf("Idle (address: %d)\n", i)
			idleComputers++
		}

		if idleComputers >= 50 {
			if nat.y == prevNat.y {
				fmt.Printf("2 consecutive packets on NAT: (x: %d, y: %d)", nat.x, nat.y)
				return
			}
			n.network[0].Input <- nat.x
			n.network[0].Input <- nat.y
			prevNat = nat
			idleComputers = 0
		}
	}
}

func main() {
	code := intcode.ReadCode("./day23/input.txt")
	network := New(code)
	network.Boot()
}

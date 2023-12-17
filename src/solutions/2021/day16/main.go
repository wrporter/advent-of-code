package main

import (
	"aoc/src/lib/go/file"
	"aoc/src/lib/go/ints"
	"aoc/src/lib/go/out"
	"aoc/src/lib/go/timeit"
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	defer timeit.Report(time.Now())

	year, day := 2021, 16
	out.Day(year, day)
	input, _ := file.ReadFile(fmt.Sprintf("./solutions/%d/day%d/input.txt", year, day))

	answer1 := part1(input)
	out.Part1(answer1)

	answer2 := part2(input)
	out.Part2(answer2)
}

func part1(input []string) interface{} {
	p := parsePacket(input[0])
	sum := sumVersions(p)
	return sum
}

// 18902248673 is too low
// 144595909277 correct
func part2(input []string) interface{} {
	p := parsePacket(input[0])
	//fmt.Println(p)
	value := calculateValue(p)
	return value
}

func parsePacket(hex string) packet {
	data := hexToBinary(hex)
	p := decode(&bits{data: data})
	return p
}

func calc(data string) int {
	p := decode(&bits{data: data})
	value := calculateValue(p)
	return value
}

func calculateValue(p packet) int {
	value := 0
	switch p.TypeID {
	case 0:
		value = 0
		for _, sub := range p.SubPackets {
			value += calculateValue(sub)
		}
	case 1:
		value = 1
		for _, sub := range p.SubPackets {
			value *= calculateValue(sub)
		}
	case 2:
		value = math.MaxInt
		for _, sub := range p.SubPackets {
			value = ints.Min(value, calculateValue(sub))
		}
	case 3:
		value = math.MinInt
		for _, sub := range p.SubPackets {
			value = ints.Max(value, calculateValue(sub))
		}
	case 4:
		value = p.LiteralValue
	case 5:
		if calculateValue(p.SubPackets[0]) > calculateValue(p.SubPackets[1]) {
			value = 1
		}
	case 6:
		if calculateValue(p.SubPackets[0]) < calculateValue(p.SubPackets[1]) {
			value = 1
		}
	case 7:
		if calculateValue(p.SubPackets[0]) == calculateValue(p.SubPackets[1]) {
			value = 1
		}
	}
	return value
}

func sumVersions(p packet) int {
	sum := p.Version
	for _, sub := range p.SubPackets {
		sum += sumVersions(sub)
	}
	return sum
}

func decode(b *bits) packet {
	p := packet{
		Binary:  b.peek(6),
		Version: b.popInt(3),
		TypeID:  b.popInt(3),
	}

	if p.TypeID == 4 {
		value := ""
		for group := ""; len(b.data) > 0 && group != "0"; {
			p.Binary += b.peek(5)
			group = b.pop(1)
			value += b.pop(4)
		}
		p.LiteralValue = binaryToInt(value)
		//fmt.Println(p)
		return p
	}

	p.Binary += b.peek(1)
	p.LengthID = b.popInt(1)

	if p.LengthID == 0 {
		p.Binary += b.peek(15)
		totalLength := b.popInt(15)

		length := 0
		for length != totalLength {
			child := decode(b)
			length += len(child.Binary)
			p.Binary += child.Binary
			p.SubPackets = append(p.SubPackets, child)
		}
	} else if p.LengthID == 1 {
		p.Binary += b.peek(11)
		numSubPackets := b.popInt(11)

		for i := 0; i < numSubPackets; i++ {
			child := decode(b)
			p.Binary += child.Binary
			p.SubPackets = append(p.SubPackets, child)
		}
	}

	//fmt.Println(p)

	return p
}

type bits struct {
	data string
}

func (b *bits) peek(amount int) string {
	return b.data[:amount]
}

func (b *bits) pop(amount int) string {
	result := b.data[:amount]
	b.data = b.data[amount:]
	return result
}

func (b *bits) popInt(amount int) int {
	return binaryToInt(b.pop(amount))
}

type packet struct {
	Binary       string   `json:"binary,omitempty"`
	Version      int      `json:"version,omitempty"`
	TypeID       int      `json:"typeId,omitempty"`
	LengthID     int      `json:"lengthId,omitempty"`
	LiteralValue int      `json:"literalValue,omitempty"`
	SubPackets   []packet `json:"subPackets,omitempty"`
}

func (p packet) String() string {
	//bytes, _ := json.Marshal(p)
	//return string(bytes)
	return fmt.Sprintf("{version: %d, typeId: %d}", p.Version, p.TypeID)
}

func hexToBinary(hex string) string {
	result := ""
	for _, char := range hex {
		result += hexCharToBinary(char)
	}
	return result
}

func hexCharToBinary(hex rune) string {
	ui, _ := strconv.ParseUint(string(hex), 16, 4)
	return fmt.Sprintf("%04b", ui)
}

func binaryToInt(binary string) int {
	value, _ := strconv.ParseInt(binary, 2, 32)
	return int(value)
}

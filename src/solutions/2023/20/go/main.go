package main

import (
	"aoc/src/lib/go/aoc"
	"aoc/src/lib/go/v2/mymath"
	"strings"
)

func part1(input string, _ ...interface{}) interface{} {
	modules := parse(input)

	countLow := 0
	countHigh := 0

	for i := 1; i <= 1000; i++ {
		signals := []Signal{button}

		for len(signals) > 0 {
			signal := signals[0]
			signals = signals[1:]

			if signal.Pulse == Low {
				countLow += 1
			} else {
				countHigh += 1
			}

			//fmt.Printf("%s\n", signal.String())
			signals = append(signals, signal.Process(modules)...)
		}
	}

	return countLow * countHigh
}

func part2(input string, _ ...interface{}) interface{} {
	modules := parse(input)

	rxInputs := getInputs(modules, "rx")
	if len(rxInputs) == 0 {
		return -1
	}
	rxInput := modules[rxInputs[0]].(*ConjunctionModule)
	first := make(map[string]int)

	for i := 1; len(first) != len(rxInput.Memory); i++ {
		signals := []Signal{button}

		for len(signals) > 0 {
			signal := signals[0]
			signals = signals[1:]

			//fmt.Printf("%s\n", signal.String())
			signals = append(signals, signal.Process(modules)...)

			for source, pulse := range rxInput.Memory {
				if _, ok := first[source]; !ok && pulse == High {
					first[source] = i
				}
			}
		}
	}

	var values []int
	for _, v := range first {
		values = append(values, v)
	}

	return mymath.LCM(values...)
}

func parse(input string) map[string]Module {
	modules := make(map[string]Module)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		destinations := strings.Split(parts[1], ", ")

		if line[0] == '%' {
			name := parts[0][1:]
			modules[name] = &FlipFlopModule{
				AbstractModule: AbstractModule{
					Name:         name,
					Destinations: destinations,
				},
			}
		} else if line[0] == '&' {
			name := parts[0][1:]
			modules[name] = &ConjunctionModule{
				AbstractModule: AbstractModule{
					Name:         name,
					Destinations: destinations,
				},
			}
		} else {
			name := parts[0]
			modules[name] = &BroadcastModule{
				AbstractModule: AbstractModule{
					Name:         name,
					Destinations: destinations,
				},
			}
		}
	}

	for _, module := range modules {
		module.Init(modules)
	}

	return modules
}

type Module interface {
	Init(modules map[string]Module)
	GetDestinations() []string
	Pulse(input Signal) []Signal
}

type AbstractModule struct {
	Name         string
	Destinations []string
}

func (m *AbstractModule) Init(_ map[string]Module) {}

func (m *AbstractModule) GetDestinations() []string {
	return m.Destinations
}

func (m *AbstractModule) Send(output Pulse) []Signal {
	signals := make([]Signal, len(m.Destinations))
	for i, destination := range m.Destinations {
		signals[i] = Signal{
			Source:      m.Name,
			Destination: destination,
			Pulse:       output,
		}
	}
	return signals
}

type FlipFlopModule struct {
	AbstractModule
	State State
}

func (m *FlipFlopModule) Init(_ map[string]Module) {
	m.State = Off
}

func (m *FlipFlopModule) Pulse(input Signal) []Signal {
	if input.Pulse == High {
		return nil
	}

	output := Low
	if m.State == Off {
		output = High
	}
	m.State ^= On

	return m.Send(output)
}

type ConjunctionModule struct {
	AbstractModule
	Memory map[string]Pulse
}

func (m *ConjunctionModule) Init(modules map[string]Module) {
	m.Memory = make(map[string]Pulse)
	for _, input := range getInputs(modules, m.Name) {
		m.Memory[input] = Low
	}
}

func (m *ConjunctionModule) Pulse(input Signal) []Signal {
	m.Memory[input.Source] = input.Pulse

	allHigh := true
	for _, pulse := range m.Memory {
		if pulse == Low {
			allHigh = false
			break
		}
	}

	output := Low
	if !allHigh {
		output = High
	}
	return m.Send(output)
}

type BroadcastModule struct {
	AbstractModule
}

func (m *BroadcastModule) Pulse(input Signal) []Signal {
	return m.Send(input.Pulse)
}

type Signal struct {
	Source      string
	Destination string
	Pulse       Pulse
}

func (s Signal) String() string {
	var sb strings.Builder
	sb.WriteString(s.Source)
	sb.WriteString(" -")
	pulse := "low"
	if s.Pulse == High {
		pulse = "high"
	}
	sb.WriteString(pulse)
	sb.WriteString("-> ")
	sb.WriteString(s.Destination)
	return sb.String()
}

func (s Signal) Process(modules map[string]Module) []Signal {
	if module, ok := modules[s.Destination]; ok {
		return module.Pulse(s)
	} else {
		return nil
	}
}

func getInputs(modules map[string]Module, destination string) []string {
	var inputs []string
	for input, module := range modules {
		for _, target := range module.GetDestinations() {
			if target == destination {
				inputs = append(inputs, input)
			}
		}
	}
	return inputs
}

type Pulse int
type State int

const (
	Low Pulse = iota
	High
)

const (
	Off State = iota
	On
)

var button = Signal{
	Source:      "button",
	Destination: "broadcaster",
	Pulse:       Low,
}

func main() {
	New().Run(nil, nil)
}

func New() aoc.Solution {
	return aoc.Solution{Year: 2023, Day: 20, Part1: part1, Part2: part2}
}

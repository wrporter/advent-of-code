package amplify

import (
	"github.com/wrporter/advent-of-code-2019/day5/public/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/arrays"
	"github.com/wrporter/advent-of-code-2019/internal/common/probability"
	"sync"
)

type AmplificationCircuit struct {
	cpu     *computer.Computer
	program []int
}

type AmplifierControllerSoftware struct {
	program      []int
	input        int
	phaseSetting int
}

type AmpCombo struct {
	MaxThrusterSignal int
	PhaseSettings     []int
}

func New(program []int) *AmplificationCircuit {
	return &AmplificationCircuit{computer.New(), program}
}

func (ac *AmplificationCircuit) Amplify(phaseSettingOptions []int) AmpCombo {
	maxCombo := AmpCombo{0, nil}

	probability.Permute(phaseSettingOptions, func(phaseSettings []int) {
		amplifiers := make([]*Amplifier, 0)
		input := make(chan int, 2)
		var wg sync.WaitGroup

		for _, phaseSetting := range phaseSettings {
			output := make(chan int, 2)
			input <- phaseSetting
			amplifier := NewAmplifier(ac.copyProgram(), input, output)
			amplifiers = append(amplifiers, amplifier)
			input = output
		}
		amplifiers[0].input <- 0
		amplifiers[len(amplifiers)-1].output = amplifiers[0].input

		for _, amplifier := range amplifiers {
			wg.Add(1)
			go ac.cpu.Thread(&wg, amplifier.program, amplifier.input, amplifier.output)
		}

		wg.Wait()
		thrusterSignal := <-amplifiers[len(amplifiers)-1].output
		if thrusterSignal > maxCombo.MaxThrusterSignal {
			maxCombo.MaxThrusterSignal = thrusterSignal
			maxCombo.PhaseSettings = arrays.CopyInts(phaseSettings)
		}
	})
	return maxCombo
}

func (ac *AmplificationCircuit) copyProgram() []int {
	return arrays.CopyInts(ac.program)
}

type Amplifier struct {
	program []int
	input   chan int
	output  chan int
}

func NewAmplifier(program []int, input chan int, output chan int) *Amplifier {
	return &Amplifier{program, input, output}
}

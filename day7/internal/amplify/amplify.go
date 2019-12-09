package amplify

import (
	"github.com/wrporter/advent-of-code-2019/day5/public/computer"
	"github.com/wrporter/advent-of-code-2019/internal/common/arrays"
	"github.com/wrporter/advent-of-code-2019/internal/common/probability"
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

func (ac *AmplificationCircuit) Amplify(numAmplifiers int, phaseSettingOptions []int) AmpCombo {
	maxCombo := AmpCombo{0, nil}
	probability.Permute(phaseSettingOptions, func(phaseSettings []int) {
		inputSignal := 0
		for _, phaseSetting := range phaseSettings {
			inputSignal = ac.cpu.Run(ac.copyProgram(), []int{phaseSetting, inputSignal})[0]
		}
		if inputSignal > maxCombo.MaxThrusterSignal {
			maxCombo.MaxThrusterSignal = inputSignal
			maxCombo.PhaseSettings = arrays.CopyInts(phaseSettings)
		}
	})
	return maxCombo
}

func (ac *AmplificationCircuit) copyProgram() []int {
	return arrays.CopyInts(ac.program)
}

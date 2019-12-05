package computer

import (
	"github.com/wrporter/advent-of-code-2019/day5/internal/convert"
	"strconv"
)

func ParseInstruction(program []int, index int) Instruction {
	return Instruction{
		Intcode: parseIntcode(program[index]),
		Address: index,
	}
}

func parseIntcode(intcode int) Intcode {
	intcodeString := strconv.Itoa(intcode)

	opCodeSpace := getOpCodeSpace(intcodeString)
	opCode := OpCode(convert.StringToInt(intcodeString[opCodeSpace:]))
	numParameters := OpCodeNumParameters[opCode]

	parameterModeCodes := convert.Reverse(intcodeString[:opCodeSpace])
	parameterModes := make([]ParameterMode, numParameters)
	for i := 0; i < len(parameterModeCodes); i++ {
		parameterModes[i] = ParameterMode(convert.RuneToInt(parameterModeCodes[i]))
	}

	return Intcode{opCode, numParameters, parameterModes}
}

func getOpCodeSpace(intcode string) int {
	space := len(intcode) - 2
	if space < 0 {
		return 0
	}
	return space
}

package computer

import (
	"github.com/wrporter/advent-of-code-2019/internal/common/conversion"
	"strconv"
)

func ParseInstruction(program *Program) Instruction {
	return Instruction{
		Intcode: parseIntcode(program.Memory[program.address]),
		Address: program.address,
	}
}

func parseIntcode(intcode int) Intcode {
	intcodeString := strconv.Itoa(intcode)

	opCodeSpace := getOpCodeSpace(intcodeString)
	opCode := OpCode(conversion.StringToInt(intcodeString[opCodeSpace:]))
	numParameters := OpCodeNumParameters[opCode]

	parameterModeCodes := conversion.Reverse(intcodeString[:opCodeSpace])
	parameterModes := make([]ParameterMode, numParameters)
	for i := 0; i < len(parameterModeCodes); i++ {
		parameterModes[i] = ParameterMode(conversion.RuneToInt(parameterModeCodes[i]))
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

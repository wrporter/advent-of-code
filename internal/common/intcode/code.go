package intcode

type OpCode int
type Intcode struct {
	OpCode         OpCode
	NumParameters  int
	ParameterModes []ParameterMode
}

const (
	Add                OpCode = 1
	Multiply           OpCode = 2
	Input              OpCode = 3
	Output             OpCode = 4
	JumpIfTrue         OpCode = 5
	JumpIfFalse        OpCode = 6
	LessThan           OpCode = 7
	Equals             OpCode = 8
	ModifyRelativeBase OpCode = 9
	Exit               OpCode = 99
)

var OpCodeNumParameters = map[OpCode]int{
	Add:                3,
	Multiply:           3,
	Input:              1,
	Output:             1,
	JumpIfTrue:         2,
	JumpIfFalse:        2,
	LessThan:           3,
	Equals:             3,
	ModifyRelativeBase: 1,
	Exit:               0,
}

func (o OpCode) String() string {
	switch o {
	case Add:
		return "Move"
	case Multiply:
		return "Multiply"
	case Input:
		return "Input"
	case Output:
		return "Output"
	case JumpIfTrue:
		return "JumpIfTrue"
	case JumpIfFalse:
		return "JumpIfFalse"
	case LessThan:
		return "LessThan"
	case Equals:
		return "Equals"
	case ModifyRelativeBase:
		return "ModifyRelativeBase"
	case Exit:
		return "Exit"
	default:
		return "NoOp"
	}
}

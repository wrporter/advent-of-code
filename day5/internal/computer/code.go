package computer

type Code int
type OpCode struct {
	Code          Code
	NumParameters int
}

type Intcode struct {
	OpCode         OpCode
	ParameterModes []ParameterMode
}

const (
	Add         Code = 1
	Multiply    Code = 2
	Input       Code = 3
	Output      Code = 4
	JumpIfTrue  Code = 5
	JumpIfFalse Code = 6
	LessThan    Code = 7
	Equals      Code = 8
	Exit        Code = 99
)

var OpCodes = map[Code]OpCode{
	Add:         {Add, 3},
	Multiply:    {Multiply, 3},
	Input:       {Input, 1},
	Output:      {Output, 1},
	JumpIfTrue:  {JumpIfTrue, 2},
	JumpIfFalse: {JumpIfFalse, 2},
	LessThan:    {LessThan, 3},
	Equals:      {Equals, 3},
	Exit:        {Exit, 0},
}

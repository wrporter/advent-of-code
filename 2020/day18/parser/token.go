package parser

import "strconv"

type (
	Token int
)

const (
	EOF Token = iota

	INT

	ADD
	SUB
	MUL
	QUO

	LPAREN
	RPAREN
)

const (
	OperatorBegin = 2
	OperatorEnd   = 5
)

var tokens = map[Token]string{
	EOF: "EOF",

	INT: "INT",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",

	LPAREN: "(",
	RPAREN: ")",
}

func (t Token) String() string {
	s := ""
	if 0 <= t && t < Token(len(tokens)) {
		s = tokens[t]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}

func (t Token) IsOperator() bool {
	return t >= OperatorBegin && t <= OperatorEnd
}

const (
	LowestPrec  = 0
	UnaryPrec   = 6
	HighestPrec = 7
)

func (t Token) Precedence() int {
	switch t {
	case ADD, SUB:
		return 4
	case MUL, QUO:
		return 5
	}
	return LowestPrec
}

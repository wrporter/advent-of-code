package parser

type (
	Interpreter struct {
	}
)

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Eval(expr Expr) int {
	switch exp := expr.(type) {
	case BinaryExpr:
		return i.EvalBinaryExpr(exp)
	case Num:
		return exp.Value
	}

	return 0
}

func (i *Interpreter) EvalBinaryExpr(exp BinaryExpr) int {
	switch exp.Op.Type {
	case ADD:
		return i.Eval(exp.Left) + i.Eval(exp.Right)
	case MUL:
		return i.Eval(exp.Left) * i.Eval(exp.Right)
	}

	return 0
}

func Evaluate(expression string) int {
	lexer := NewLexer(expression)
	p := NewParser(lexer)
	interpreter := NewInterpreter()
	tree := p.Parse()

	return interpreter.Eval(tree)
}

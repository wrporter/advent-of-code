package parser

import "fmt"

type (
	Expr interface {
	}

	BinaryExpr struct {
		Left  Expr
		Op    LexedToken
		Right Expr
	}

	ParenExpr struct {
	}

	Num struct {
		Token LexedToken
		Value int
	}

	Parser struct {
		lexer        *Lexer
		currentToken LexedToken
	}
)

var operators = map[Token]bool{
	ADD: true,
	MUL: true,
}

func NewNum(token LexedToken) Num {
	return Num{
		Token: token,
		Value: token.Value.(int),
	}
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		lexer:        lexer,
		currentToken: lexer.getNextToken(),
	}
}

func (p *Parser) eat(tokenType Token) {
	if p.currentToken.Type == tokenType {
		p.currentToken = p.lexer.getNextToken()
	} else {
		fmt.Printf("Error parsing token: %s\n", tokenType.String())
	}
}

func (p *Parser) factor() Expr {
	token := p.currentToken
	if token.Type == INT {
		p.eat(INT)
		return NewNum(token)
	} else if token.Type == LPAREN {
		p.eat(LPAREN)
		node := p.expr()
		p.eat(RPAREN)
		return node
	}
	return nil
}

func (p *Parser) term() Expr {
	node := p.factor()

	for operators[p.currentToken.Type] {
		token := p.currentToken
		p.eat(p.currentToken.Type)
		node = BinaryExpr{
			Left:  node,
			Op:    token,
			Right: p.factor(),
		}
	}

	return node
}

func (p *Parser) expr() Expr {
	return p.term()
}

func (p *Parser) Parse() Expr {
	return p.expr()
}

func ParseExpr(expr string) Expr {
	lexer := NewLexer(expr)
	parser := NewParser(lexer)
	return parser.Parse()
}

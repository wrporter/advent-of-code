package parser

import (
	"strconv"
	"unicode"
)

type (
	Lexer struct {
		text        string
		pos         int
		currentChar string
	}

	LexedToken struct {
		Type  Token
		Value interface{}
	}
)

func NewNumLexedToken(value int) LexedToken {
	return LexedToken{INT, value}
}

func NewLexer(text string) *Lexer {
	return &Lexer{
		text:        text,
		pos:         0,
		currentChar: string(text[0]),
	}
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos > len(l.text)-1 {
		l.currentChar = ""
	} else {
		l.currentChar = string(l.text[l.pos])
	}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.currentChar[0])) {
		l.advance()
	}
}

func (l *Lexer) integer() int {
	result := ""
	for l.currentChar != "" && unicode.IsDigit(rune(l.currentChar[0])) {
		result += l.currentChar
		l.advance()
	}
	value, _ := strconv.ParseInt(result, 10, 64)
	return int(value)
}

func (l *Lexer) getNextToken() LexedToken {
	for l.currentChar != "" {
		if unicode.IsSpace(rune(l.currentChar[0])) {
			l.skipWhitespace()
			continue
		} else if unicode.IsDigit(rune(l.currentChar[0])) {
			return LexedToken{INT, l.integer()}
		} else if l.currentChar == tokens[ADD] {
			l.advance()
			return LexedToken{Type: ADD}
		} else if l.currentChar == tokens[MUL] {
			l.advance()
			return LexedToken{Type: MUL}
		} else if l.currentChar == tokens[LPAREN] {
			l.advance()
			return LexedToken{Type: LPAREN}
		} else if l.currentChar == tokens[RPAREN] {
			l.advance()
			return LexedToken{Type: RPAREN}
		}
	}

	return LexedToken{Type: EOF}
}

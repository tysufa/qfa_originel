package lexer

import "github.com/tysufa/qfa/token"

// "github.com/tysufa/qfa/token"

type Lexer struct {
	curChar  byte
	peekChar byte
	input    string
	pos      int
	line     int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, pos: -1, line: 1}
	l.nextChar()
	l.nextChar()

	return l
}

func (l *Lexer) getToken() token.Token {
	var tok token.Token

	switch l.curChar {
	case '!':
		tok.Type = token.BANG
		tok.Value = string(l.curChar)
	case '/':
		tok.Type = token.SLASH
		tok.Value = string(l.curChar)
	case '*':
		tok.Type = token.STAR
		tok.Value = string(l.curChar)
	case '-':
		tok.Type = token.MINUS
		tok.Value = string(l.curChar)
	case '+':
		tok.Type = token.PLUS
		tok.Value = string(l.curChar)
	case '(':
		tok.Type = token.LPAR
		tok.Value = string(l.curChar)
	case ')':
		tok.Type = token.RPAR
		tok.Value = string(l.curChar)
	case '{':
		tok.Type = token.LBR
		tok.Value = string(l.curChar)
	case '}':
		tok.Type = token.RBR
		tok.Value = string(l.curChar)
	}
	l.nextChar()
	return tok
}

func (l *Lexer) nextChar() {
	l.curChar = l.peekChar
	if l.pos+1 < len(l.input) {
		l.pos++
		l.peekChar = l.getCurChar()
	}
}

func (l *Lexer) getCurChar() byte {
	return l.input[l.pos]
}

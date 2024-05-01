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

	l.skipSpace()
	switch l.curChar {
	case 0:
		tok.Type = token.EOF
		tok.Value = ""
	case ';':
		tok.Type = token.SEMICOLON
		tok.Value = string(l.curChar)
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
	case '=':
		tok.Type = token.EQ
		tok.Value = string(l.curChar)
	default:
		if l.isLetter() {
			literal := l.getWord()
			if literal == "let" {
				tok.Type = token.LET
			} else {
				tok.Type = token.IDENT
			}
			tok.Value = literal
		}
	}

	l.nextChar()
	print(string(l.curChar), "\n")
	return tok
}

func (l *Lexer) getWord() string {
	res := ""
	for l.isLetter() {
		res += string(l.curChar)
		l.nextChar()
	}

	return res
}

func (l *Lexer) isLetter() bool {
	return ('a' <= l.curChar && l.curChar >= 'z') || ('A' <= l.curChar && l.curChar >= 'Z')
}

func (l *Lexer) skipSpace() {
	for l.curChar == ' ' {
		l.nextChar()
	}
}

func (l *Lexer) nextChar() {
	l.curChar = l.peekChar
	if l.pos+1 < len(l.input) {
		l.pos++
		l.peekChar = l.getCurChar()
	} else {
		l.peekChar = 0
	}
}

func (l *Lexer) getCurChar() byte {
	return l.input[l.pos]
}

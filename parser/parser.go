package parser

import (
	"fmt"
	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
)

type Parser struct {
	lex       lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l lexer.Lexer) *Parser {
	cur := l.GetToken()
	peek := l.GetToken()
	return &Parser{lex: l, curToken: cur, peekToken: peek}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.GetToken()
}

func (p *Parser) GetStatements() ast.Program {
	res := ast.Program{}

	for p.curToken.Type != token.EOF {
		switch p.curToken.Type {
		case token.LET:
			stmt := p.parseLet()
			if stmt != nil {
				res.Statements = append(res.Statements, stmt)
			}
		}
		p.nextToken()
	}
	return res
}

func (p *Parser) parseLet() *ast.LetStatement {
	let := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	let.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
	if !p.expectPeek(token.EQ) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return let
}

func (p *Parser) expectPeek(expectToken token.TokenType) bool {
	if p.peekToken.Type == expectToken {
		p.nextToken()
		return true
	} else {
		err := fmt.Sprintf("expected '%v', got '%v' instead at line %v", expectToken, p.peekToken.Type, p.curToken.Line)
		p.errors = append(p.errors, err)
		p.nextToken()
		return false
	}
}

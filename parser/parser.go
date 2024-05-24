package parser

import (
	"fmt"
	"strconv"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/token"
)

type (
	prefixParseFn func() ast.Expression
)

type Parser struct {
	lex            lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
}

func New(l lexer.Lexer) *Parser {
	cur := l.GetToken()
	peek := l.GetToken()
	p := &Parser{lex: l, curToken: cur, peekToken: peek}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)

	p.prefixParseFns[token.IDENT] = p.parseIdent
	p.prefixParseFns[token.INT] = p.parseIntegerLiteral
	p.prefixParseFns[token.MINUS] = p.parsePrefixExpression
	p.prefixParseFns[token.BANG] = p.parsePrefixExpression

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.GetToken()
}

func (p *Parser) GetStatements() ast.Program {
	res := ast.Program{}

	for p.curToken.Type != token.EOF {
		var stmt ast.Statement
		switch p.curToken.Type {
		case token.LET:
			stmt = p.parseLet()
		default:
			stmt = p.parseExpressionStatement()
		}

		if stmt != nil {
			res.Statements = append(res.Statements, stmt)
		}
		p.nextToken()
	}
	return res
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	prefix := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Value}
	p.nextToken()
	prefix.Right = p.parseExpression()
	return prefix
}

func (p *Parser) parseIdent() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, err := strconv.Atoi(p.curToken.Value)
	if err != nil {
		err := fmt.Sprintf("could not convert %s to an integer", p.curToken.Value)
		p.errors = append(p.errors, err)
		return nil
	} else {
		return &ast.IntegerLiteral{Token: p.curToken, Value: val}
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression()

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
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

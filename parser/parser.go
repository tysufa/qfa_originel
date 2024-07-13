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
	infixParseFn  func(left ast.Expression) ast.Expression
)

const (
	_ = iota
	LOWEST
	EQUAL
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Parser struct {
	lex            lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	Errors         []string
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l lexer.Lexer) *Parser {
	cur := l.GetToken()
	peek := l.GetToken()
	p := &Parser{lex: l, curToken: cur, peekToken: peek}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)

	p.prefixParseFns[token.IDENT] = p.parseIdent
	p.prefixParseFns[token.TRUE] = p.parseBool
	p.prefixParseFns[token.FALSE] = p.parseBool
	p.prefixParseFns[token.INT] = p.parseIntegerLiteral
	p.prefixParseFns[token.MINUS] = p.parsePrefixExpression
	p.prefixParseFns[token.BANG] = p.parsePrefixExpression
	p.prefixParseFns[token.LPAR] = p.parseGroupExpression
	p.prefixParseFns[token.IF] = p.parseIfExpression
	p.prefixParseFns[token.FN] = p.parseFunctionLiteral

	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.infixParseFns[token.PLUS] = p.parseInfixExpression
	p.infixParseFns[token.MINUS] = p.parseInfixExpression
	p.infixParseFns[token.STAR] = p.parseInfixExpression
	p.infixParseFns[token.SLASH] = p.parseInfixExpression
	p.infixParseFns[token.EQEQ] = p.parseInfixExpression
	p.infixParseFns[token.NEQ] = p.parseInfixExpression
	p.infixParseFns[token.LT] = p.parseInfixExpression
	p.infixParseFns[token.GT] = p.parseInfixExpression
	p.infixParseFns[token.GEQT] = p.parseInfixExpression
	p.infixParseFns[token.LEQT] = p.parseInfixExpression
	p.infixParseFns[token.LPAR] = p.parseCallExpression

	return p
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}
	if p.peekToken.Type == token.RPAR {
		p.nextToken()
		return args
	}
	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.peekToken.Type == token.COMMA {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(token.RPAR) {
		return nil
	}
	return args
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	fn := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAR) {
		return nil
	}
	fn.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBR) {
		return nil
	}

	fn.Body = *p.parseBlockStatement()

	return fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	idents := []*ast.Identifier{}

	if p.peekToken.Type == token.RPAR {
		p.nextToken()
		return idents
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
	idents = append(idents, ident)

	for p.peekToken.Type == token.COMMA {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
		idents = append(idents, ident)
	}

	if !p.expectPeek(token.RPAR) {
		return nil
	}

	return idents
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.GetToken()
}

func (p *Parser) parseStatement() ast.Statement {
	var stmt ast.Statement
	switch p.curToken.Type {
	case token.LET:
		stmt = p.parseLet()
	case token.RETURN:
		stmt = p.parseReturn()
  case token.IDENT:
    if p.peekToken.Type == token.EQ{
      stmt = p.parseAssignement()
    } else{
      stmt = p.parseExpressionStatement()
    }
    
	default:
		stmt = p.parseExpressionStatement()
	}
	return stmt
}

func (p *Parser) GetStatements() ast.Program {
	res := ast.Program{}

	for p.curToken.Type != token.EOF {
		var stmt ast.Statement
		stmt = p.parseStatement()

		if stmt != nil {
			res.Statements = append(res.Statements, stmt)
		}
		p.nextToken()
	}
	return res
}

func (p *Parser) parseAssignement() *ast.AssignementStatement{
	ass := &ast.AssignementStatement{Token: p.curToken}

	ass.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
	p.nextToken()
	p.nextToken()

	ass.Value = p.parseExpression(LOWEST)

  if p.curToken.Type == token.SEMICOLON{
    p.nextToken()
  }

	return ass
}

var precedences = map[token.TokenType]int{
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.STAR:  PRODUCT,
	token.SLASH: PRODUCT,
	token.GT:    LESSGREATER,
	token.LT:    LESSGREATER,
	token.LEQT:  LESSGREATER,
	token.GEQT:  LESSGREATER,
	token.EQEQ:  EQUAL,
	token.NEQ:   EQUAL,
	token.LPAR:  CALL,
}

func (p *Parser) getPeekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) getPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	prefix := &ast.PrefixExpression{Token: p.curToken, Operator: p.curToken.Value}
	p.nextToken()
	prefix.Right = p.parseExpression(PREFIX)
	return prefix
}

func (p *Parser) parseBool() ast.Expression {
	if p.curToken.Type == token.TRUE {
		return &ast.Boolean{Token: p.curToken, Value: true}
	} else {
		return &ast.Boolean{Token: p.curToken, Value: false}
	}

}

func (p *Parser) parseIdent() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Value}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	val, err := strconv.Atoi(p.curToken.Value)
	if err != nil {
		err := fmt.Sprintf("could not convert %s to an integer", p.curToken.Value)
		p.Errors = append(p.Errors, err)
		return nil
	} else {
		return &ast.IntegerLiteral{Token: p.curToken, Value: val}
	}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	infix := &ast.InfixExpression{Token: p.curToken, Left: left, Operator: p.curToken.Value}
	precedence := p.getPrecedence()
	p.nextToken()
	infix.Right = p.parseExpression(precedence)
	return infix
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.Errors = append(p.Errors, fmt.Sprintf("no parse function found for prefix %v", p.curToken.Type))
		return nil
	}

	leftExp := prefix()

	for p.peekToken.Type != token.SEMICOLON && precedence < p.getPeekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseGroupExpression() ast.Expression {
	p.nextToken()

	expr := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAR) {
		return nil
	}

	return expr
}

func (p *Parser) parseReturn() *ast.ReturnStatement {
	ret := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	ret.Value = p.parseExpression(LOWEST)

	p.expectPeek(token.SEMICOLON)

	return ret
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
	p.nextToken()

	let.Value = p.parseExpression(LOWEST)

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return let
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for p.curToken.Type != token.RBR && p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block

}

func (p *Parser) parseIfExpression() ast.Expression {
	is := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAR) {
		return nil
	}
	p.nextToken()

	is.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAR) {
		return nil
	}

	if !p.expectPeek(token.LBR) {
		return nil
	}

	is.Consequences = p.parseBlockStatement()
	if p.peekToken.Type == token.ELSE {
		p.nextToken()
		if !p.expectPeek(token.LBR) {
			return nil
		}
		is.ElseConsequences = p.parseBlockStatement()
	} else {
		is.ElseConsequences = &ast.BlockStatement{}
	}

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return is
}

func (p *Parser) expectPeek(expectToken token.TokenType) bool {
	if p.peekToken.Type == expectToken {
		p.nextToken()
		return true
	} else {
		err := fmt.Sprintf("expected '%v', got '%v' instead at line %v", expectToken, p.peekToken.Type, p.curToken.Line)
		p.Errors = append(p.Errors, err)
		p.nextToken()
		return false
	}
}

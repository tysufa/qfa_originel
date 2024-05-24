package ast

import (
	"github.com/tysufa/qfa/token"
)

type Program struct {
	Statements []Statement
}

type Statement interface {
	TokenLiteral() string
	StatementNode()
}

type Expression interface {
	TokenLiteral() string
	ExpressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Value }
func (i *Identifier) ExpressionNode()      {}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Value }
func (es *ExpressionStatement) StatementNode()       {}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (ps *PrefixExpression) TokenLiteral() string { return ps.Token.Value }
func (ps *PrefixExpression) ExpressionNode()      {}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Value }
func (il *IntegerLiteral) ExpressionNode()      {}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }
func (ls *LetStatement) StatementNode()       {}

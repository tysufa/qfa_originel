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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }
func (ls *LetStatement) StatementNode()       {}

package ast

import (
	"bytes"
	"strings"

	"github.com/tysufa/qfa/token"
)

type Node interface {
	String() string
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

type Statement interface {
	Node
	TokenLiteral() string
	StatementNode()
}

type Expression interface {
	Node
	TokenLiteral() string
	ExpressionNode()
}

type CallExpression struct {
	Token     token.Token // ( token
	Function  Expression  //identifier ou FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) TokenLiteral() string { return ce.TokenLiteral() }
func (ce *CallExpression) ExpressionNode()      {}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.TokenLiteral() }
func (fl *FunctionLiteral) ExpressionNode()      {}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Value }
func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) String() string       { return "return " + rs.Value.String() }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Value }
func (i *Identifier) ExpressionNode()      {}
func (i *Identifier) String() string       { return i.Value }

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Value }
func (es *ExpressionStatement) StatementNode()       {}
func (es *ExpressionStatement) String() string {
	return es.Expression.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (is *InfixExpression) TokenLiteral() string { return is.Token.Value }
func (is *InfixExpression) ExpressionNode()      {}
func (is *InfixExpression) String() string {
	res := "(" + is.Left.String() + is.Operator + is.Right.String() + ")"
	return res
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (ps *PrefixExpression) TokenLiteral() string { return ps.Token.Value }
func (ps *PrefixExpression) ExpressionNode()      {}
func (ps *PrefixExpression) String() string {
	res := "(" + ps.Operator + ps.Right.String() + ")"
	return res
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string { return b.Token.Value }
func (b *Boolean) ExpressionNode()      {}
func (b *Boolean) String() string {
	return b.Token.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Value }
func (il *IntegerLiteral) ExpressionNode()      {}
func (il *IntegerLiteral) String() string {
	return il.Token.Value
}

type BlockStatement struct {
	Token      token.Token // { token
	Statements []Statement
}

func (sb *BlockStatement) TokenLiteral() string { return sb.Token.Value }
func (sb *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("")
	for _, stmt := range sb.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

type IfExpression struct {
	Token            token.Token
	Condition        Expression
	Consequences     *BlockStatement
	ElseConsequences *BlockStatement
}

func (is *IfExpression) TokenLiteral() string { return is.Token.Value }
func (is *IfExpression) ExpressionNode()      {}
func (is *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(is.Condition.String())
	out.WriteString("{")
	out.WriteString(is.Consequences.String())
	out.WriteString("}")
	out.WriteString("else")
	out.WriteString("{")
	out.WriteString(is.ElseConsequences.String())
	out.WriteString("}")
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Value }
func (ls *LetStatement) StatementNode()       {}
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString("let " + ls.Name.String() + " = " + ls.Value.String() + ";")

	return out.String()
}

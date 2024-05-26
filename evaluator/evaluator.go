package evaluator

import (
	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Evaluate(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.ExpressionStatement:
		return Evaluate(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return boolToBoolObject(node.Value)
	case *ast.InfixExpression:
		return evaluteExression(node)
	}

	return nil
}

func boolToBoolObject(b bool) *object.Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func evaluteExression(node *ast.InfixExpression) object.Object {
	res := &object.Integer{}
	rightInt, ok := Evaluate(node.Right).(*object.Integer)
	if !ok {
		return nil
	}
	right := rightInt.Value
	leftInt, ok := Evaluate(node.Left).(*object.Integer)
	if !ok {
		return nil
	}
	left := leftInt.Value

	switch node.Operator {
	case "+":
		return &object.Integer{Value: right + left}
	case "-":
		return &object.Integer{Value: right - left}
	case "*":
		return &object.Integer{Value: right * left}
	case "/":
		return &object.Integer{Value: right / left}
	}

	return res
}

func EvaluateStatements(program ast.Program) []object.Object {
	var res []object.Object
	for _, stmt := range program.Statements {
		res = append(res, Evaluate(stmt))
	}

	return res
}

package evaluator

import (
	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/object"
)

func Evaluate(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.ExpressionStatement:
		return Evaluate(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.InfixExpression:
		right := Evaluate(node.Right).(*object.Integer).Value
		left := Evaluate(node.Left).(*object.Integer).Value
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
	}

	return nil
}

func EvaluateStatements(program ast.Program) []object.Object {
	var res []object.Object
	for _, stmt := range program.Statements {
		res = append(res, Evaluate(stmt))
	}

	return res
}

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
	case *ast.PrefixExpression:
		return evaluatePrefix(node)
	case *ast.InfixExpression:
		return evaluteInfixExpression(node)
	}

	return nil
}

func evaluatePrefix(node *ast.PrefixExpression) object.Object {
	right := Evaluate(node.Right)
	if node.Operator == "!" {
		return evaluateBangOperatorExpression(right)
	} else if node.Operator == "-" {
		return evaluateMinusOperatorExpression(right)
	} else {
		return nil
	}
}

func evaluateMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return nil
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evaluateBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return nil
	}
}

func boolToBoolObject(b bool) *object.Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func evaluteInfixExpression(node *ast.InfixExpression) object.Object {
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
	case "==":
		return boolToBoolObject(right == left)
	case "!=":
		return boolToBoolObject(right != left)
	case ">":
		return boolToBoolObject(right > left)
	case ">=":
		return boolToBoolObject(right >= left)
	case "<":
		return boolToBoolObject(right < left)
	case "<=":
		return boolToBoolObject(right <= left)
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

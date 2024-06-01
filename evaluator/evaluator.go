package evaluator

import (
	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func EvaluateProgram(statements []ast.Statement) *object.BlockObject {
	var res object.BlockObject
	for _, stmt := range statements {
		blockEval := Evaluate(stmt)
		blockStmt, ok := blockEval.(*object.BlockObject)
		if ok {
			for _, stmt := range blockStmt.Block {
				res.Block = append(res.Block, stmt)
			}
		} else {
			res.Block = append(res.Block, Evaluate(stmt))
		}

	}

	return &res
}

func EvaluateBlockStatement(block *ast.BlockStatement) object.Object {
	var res object.BlockObject
	for _, stmt := range block.Statements {
		res.Block = append(res.Block, Evaluate(stmt))
	}

	return &res
}

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
	case *ast.IfExpression:
		return evaluateIfExpression(node)
	case *ast.InfixExpression:
		return evaluateInfixExpression(node)
	case *ast.BlockStatement:
		return EvaluateBlockStatement(node)
	}

	return nil
}

func evaluateIfExpression(node *ast.IfExpression) object.Object {
	resCondition := Evaluate(node.Condition)
	cond := boolToBoolObject(resCondition.Inspect() == "true")

	if cond.Value {
		return EvaluateProgram(node.Consequences.Statements)
	} else {
		consequences := EvaluateProgram(node.ElseConsequences.Statements)
		if consequences.Block == nil {
			return nil
		}
		return consequences
	}
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
		return NULL
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
		return NULL
	}
}

func boolToBoolObject(b bool) *object.Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func evaluateIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "==":
		return boolToBoolObject(leftVal == rightVal)
	case "!=":
		return boolToBoolObject(leftVal != rightVal)
	case ">":
		return boolToBoolObject(leftVal > rightVal)
	case ">=":
		return boolToBoolObject(leftVal >= rightVal)
	case "<":
		return boolToBoolObject(leftVal < rightVal)
	case "<=":
		return boolToBoolObject(leftVal <= rightVal)

	default:
		return NULL
	}
}

func evaluateInfixExpression(node *ast.InfixExpression) object.Object {
	left := Evaluate(node.Left)
	right := Evaluate(node.Right)

	switch {
	case left.Type() != right.Type():
		return &object.Null{}
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evaluateIntegerInfixExpression(node.Operator, left, right)
	case node.Operator == "==":
		return boolToBoolObject(left == right)
	case node.Operator == "!=":
		return boolToBoolObject(left != right)
	}

	return NULL
}

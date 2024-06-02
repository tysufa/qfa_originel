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

func EvaluateProgram(statements []ast.Statement) []object.Object {
	var program []object.Object

	for _, stmt := range statements {
		stmtVal := Evaluate(stmt)

		switch stmtVal := stmtVal.(type) {
		case *object.Return:
			program = append(program, Evaluate(stmt))
			return program
		case *object.BlockObject:
			if stmtVal.Return {
				program = append(program, stmtVal)
				return program
			} else {
				program = append(program, stmtVal)
			}
		default:
			program = append(program, Evaluate(stmt))
		}

	}

	return program
}

func EvaluateBlockStatement(block *ast.BlockStatement) object.Object {
	res := object.BlockObject{Return: false}
	for _, stmt := range block.Statements {
		stmtVal := Evaluate(stmt)

		switch stmtVal := stmtVal.(type) {
		case *object.Return:
			res.Block = append(res.Block, Evaluate(stmt))
			res.Return = true
			return &res
		case *object.BlockObject:
			if stmtVal.Return {
				res.Block = append(res.Block, stmtVal)
				return &res
			} else {
				res.Block = append(res.Block, stmtVal)
			}
		default:
			res.Block = append(res.Block, Evaluate(stmt))
		}

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
	case *ast.ReturnStatement:
		return &object.Return{Value: Evaluate(node.Value)}
	}

	return nil
}

func evaluateIfExpression(node *ast.IfExpression) object.Object {
	resCondition := Evaluate(node.Condition)
	cond := boolToBoolObject(resCondition == TRUE)

	if cond.Value {
		return EvaluateBlockStatement(node.Consequences)
	} else {
		consequences := EvaluateBlockStatement(node.ElseConsequences)
		if node.ElseConsequences == nil {
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

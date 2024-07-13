package evaluator

import (
	"fmt"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func EvaluateProgram(statements []ast.Statement, env *object.Environment) []object.Object {
	var program []object.Object

	for _, stmt := range statements {
		stmtVal := Evaluate(stmt, env)

		switch stmtVal := stmtVal.(type) {
		case *object.Return:
			program = append(program, Evaluate(stmt, env))
			return program
		case *object.BlockObject:
			if stmtVal.Return {
				program = append(program, stmtVal)
				return program
			} else {
				program = append(program, stmtVal)
			}
		case *object.Error:
			program = append(program, stmtVal)
			return program
		default:
			program = append(program, stmtVal)
		}

	}

	return program
}

func EvaluateBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	res := object.BlockObject{Return: false}
	for _, stmt := range block.Statements {
		stmtVal := Evaluate(stmt, env)

		switch stmtVal := stmtVal.(type) {
		case *object.Return:
			res.Block = append(res.Block, stmtVal)
			res.Return = true
			return &res
		case *object.BlockObject:
			if stmtVal.Return {
				res.Block = append(res.Block, stmtVal)
				res.Return = true
				return &res
			} else {
				res.Block = append(res.Block, stmtVal)
			}
		case *object.Error:
			res.Block = append(res.Block, stmtVal)
			res.Return = true
			return &res
		default:
			res.Block = append(res.Block, Evaluate(stmt, env))
		}

	}

	return &res
}

func Evaluate(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.LetStatement:
		val := Evaluate(node.Value, env)
		env.Set(node.Name.Value, val)
	case *ast.AssignementStatement:
    val := Evaluate(node.Value, env)
    env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.ExpressionStatement:
		return Evaluate(node.Expression, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return boolToBoolObject(node.Value)
	case *ast.PrefixExpression:
		return evaluatePrefix(node, env)
	case *ast.IfExpression:
		return evaluateIfExpression(node, env)
	case *ast.InfixExpression:
		return evaluateInfixExpression(node, env)
	case *ast.BlockStatement:
		return EvaluateBlockStatement(node, env)
	case *ast.ReturnStatement:
		return &object.Return{Value: Evaluate(node.Value, env)}
	}

	return nil
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newErr("identifier not found: %v", node.Value)
	}
	return val
}

func evaluateIfExpression(node *ast.IfExpression, env *object.Environment) object.Object {
	resCondition := Evaluate(node.Condition, env)
	cond := boolToBoolObject(resCondition == TRUE)

	if cond.Value {
		return EvaluateBlockStatement(node.Consequences, env)
	} else {
		consequences := EvaluateBlockStatement(node.ElseConsequences, env)
		if node.ElseConsequences == nil {
			return nil
		}
		return consequences
	}
}

func evaluatePrefix(node *ast.PrefixExpression, env *object.Environment) object.Object {
	right := Evaluate(node.Right, env)
	if node.Operator == "!" {
		return evaluateBangOperatorExpression(right)
	} else if node.Operator == "-" {
		return evaluateMinusOperatorExpression(right)
	} else {
		return newErr("unknown operator: %s", node.Operator)
		return nil
	}
}

func evaluateMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newErr("unknown operator: -%v", right.Type())
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
		return newErr("unknown operator : !%v", right.Type())
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

func evaluateInfixExpression(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := Evaluate(node.Left, env)
	right := Evaluate(node.Right, env)

	switch {
	case left.Type() != right.Type():
		return newErr("type mismatch: %s%s%s", left.Type(), node.Operator, right.Type())
		return &object.Null{}
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evaluateIntegerInfixExpression(node.Operator, left, right)
	case node.Operator == "==":
		return boolToBoolObject(left == right)
	case node.Operator == "!=":
		return boolToBoolObject(left != right)
	default:
		return newErr("unknown operator: %v%v%v", left.Type(), node.Operator, right.Type())
	}

	return NULL
}

func newErr(format string, a ...interface{}) object.Object {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

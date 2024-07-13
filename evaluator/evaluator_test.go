package evaluator

import (
	"testing"

	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/object"
	"github.com/tysufa/qfa/parser"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; a = a+1; a;", 6},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
    // t.Fatalf("%v", testEval(tt.input)[1].Inspect())
    evaluationRes := testEval(tt.input)
    for _, res := range evaluationRes{
      if res != nil{
        testIntegerObject(t, res, tt.expected)
      }
    }
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER+BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER+BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN+BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN+BOOLEAN",
		},
		// {
		// 	"if (10 > 1) { true + false; }",
		// 	"unknown operator: BOOLEAN + BOOLEAN",
		// },
		{
			"foobar",
			"identifier not found: foobar",
		},
		// 		{
		// 			`
		// 			if (10 > 1) {
		// 				if (10 > 1) {
		// 					return true + false;
		// 				}
		// 			return 1;
		// 			}
		// `,
		// 			"unknown operator: BOOLEAN + BOOLEAN",
		// 		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated[len(evaluated)-1].(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// {"if (1+2 == 4){3*2} else {3+2*4}", 11},
		// {"if (true){if (true) {return 1;} 5} else {3+2*4}", 1},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated[0], tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, res int) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T", obj)
	}
	if result.Value != res {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, res)
	}
}

func TestIntegerEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50}}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated[0], tt.expected)
	}
}

func testEval(input string) []object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.GetStatements()

	env := object.NewEnvironment()

	return EvaluateProgram(program.Statements, env)
}

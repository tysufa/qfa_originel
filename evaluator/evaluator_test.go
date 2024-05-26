package evaluator

import (
	"testing"

	"github.com/tysufa/qfa/lexer"
	"github.com/tysufa/qfa/object"
	"github.com/tysufa/qfa/parser"
)

func TestIntegerEvaluation(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated[0], tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, res int) {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
	}
	if result.Value != res {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, res)
	}
}

func testEval(input string) []object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.GetStatements()

	return EvaluateStatements(program)
}

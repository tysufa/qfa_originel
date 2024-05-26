package parser

import (
	"testing"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
)

func TestPriorityOperations(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult string
	}{
		{"1 * 2 + 3", "((1*2)+3)"},
		{"1 / 2 - 3", "((1/2)-3)"},
		{"1 * 2 + 3 / 4 - 5", "(((1*2)+(3/4))-5)"},
		{"-1 * 2 + !3", "(((-1)*2)+(!3))"},
		{"-a*b", "((-a)*b)"},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a+b)+c)",
		},
		{
			"a + b - c",
			"((a+b)-c)",
		},
		{
			"a * b * c",
			"((a*b)*c)",
		},
		{
			"a * b / c",
			"((a*b)/c)",
		},
		{
			"a + b / c",
			"(a+(b/c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a+(b*c))+(d/e))-f)",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3+(4*5))==((3*1)+(4*5)))",
		},
		{
			"1 + (2 + 3) + 4",
			"((1+(2+3))+4)",
		},
		{
			"(5 + 5) * 2",
			"((5+5)*2)",
		},
		{
			"2 / (5 + 5)",
			"(2/(5+5))",
		},
		{
			"-(5 + 5)",
			"(-(5+5))",
		},
		{
			"!(true == true)",
			"(!(true==true))",
		},
		{
			"1 + 2 < 4 == 5 - 4 >= 1",
			"(((1+2)<4)==((5-4)>=1))",
		},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		if len(p.errors) > 0 {
			for _, err := range p.errors {
				t.Errorf(err)
			}
			t.FailNow()
		}

		exprStmt, ok := stmts.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("expected ExpressionStatement got %T instead", stmts.Statements[0])
		}

		if exprStmt.String() != test.expectedResult {
			t.Fatalf("expected %s, but got %s instead", test.expectedResult, exprStmt.String())
		}

	}
}

func TestInfixExpressions(t *testing.T) {
	tests := []struct {
		input            string
		leftValue        int
		expectedOperator string
		rightValue       int
	}{
		{"5 + 3;", 5, "+", 3},
		{"5 - 3;", 5, "-", 3},
		{"5 * 3;", 5, "*", 3},
		{"5 / 3;", 5, "/", 3},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		if len(p.errors) > 0 {
			for _, err := range p.errors {
				t.Errorf(err)
			}
			t.FailNow()
		}

		if len(stmts.Statements) != 1 {
			t.Fatalf("wrong number of stmts.Statements, expected 1, got %v instead", len(stmts.Statements))
		}

		exprStmt, ok := stmts.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("expected ExpressionStatement got %T instead", stmts.Statements[0])
		}

		infix, ok := exprStmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("expected PrefixExpression, got %T instead", exprStmt.Expression)
		}
		if infix.Operator != test.expectedOperator {
			t.Fatalf("Wrong operator, expected %s, got %s instead", test.expectedOperator, infix.Operator)
		}
		testIntegerLiteral(t, infix.Left, test.leftValue)
		testIntegerLiteral(t, infix.Right, test.rightValue)
	}
}

func TestPrefixExpressions(t *testing.T) {
	tests := []struct {
		input            string
		expectedOperator string
		expectedValue    int
	}{
		{"-5;", "-", 5},
		{"!42", "!", 42},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		if len(p.errors) > 0 {
			for _, err := range p.errors {
				t.Errorf(err)
			}
			t.FailNow()
		}

		if len(stmts.Statements) != 1 {
			t.Fatalf("wrong number of stmts.Statements, expected 1, got %v instead", len(stmts.Statements))
		}

		exprStmt, ok := stmts.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("expected ExpressionStatement got %T instead", stmts.Statements[0])
		}

		prefix, ok := exprStmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expected PrefixExpression, got %T instead", exprStmt.Expression)
		}
		if prefix.Operator != test.expectedOperator {
			t.Fatalf("Wrong operator, expected %s, got %s instead", test.expectedOperator, prefix.Operator)
		}
		testIntegerLiteral(t, prefix.Right, test.expectedValue)
	}
}

func testIntegerLiteral(t *testing.T, integerExpression ast.Expression, expectedValue int) {
	integ, ok := integerExpression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expected IntegerLiteral got %T instead", integ)
	}

	if integ.Value != expectedValue {
		t.Fatalf("wrong integer value, expected %v, got %v instead", expectedValue, integ.Value)
	}

}

func TestIntegerLiteralExpressions(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	stmt := p.GetStatements()

	if len(stmt.Statements) != 1 {
		t.Fatalf("wrong statements number, expected 1 got %v instead", len(stmt.Statements))
	}

	exprStmt, ok := stmt.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("expected ExpressionStatement got %T instead", stmt.Statements[0])
	}

	integ, ok := exprStmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expected ExpressionStatement got %T instead", exprStmt.Expression)
	}

	if integ.Value != 5 {
		t.Fatalf("expected ident of value foo, got %v instead", integ.Value)
	}
}

func TestIdentExpressions(t *testing.T) {
	input := "foo;"

	l := lexer.New(input)
	p := New(l)

	stmt := p.GetStatements()

	if len(stmt.Statements) != 1 {
		t.Fatalf("wrong statements number, expected 1 got %v instead", len(stmt.Statements))
	}

	exprStmt, ok := stmt.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("expected ExpressionStatement got %T instead", stmt.Statements[0])
	}

	ident, ok := exprStmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("expected ExpressionStatement got %T instead", exprStmt.Expression)
	}

	if ident.Value != "foo" {
		t.Fatalf("expected ident of value foo, got %s instead", ident.Value)
	}

}

func TestLetStatements(t *testing.T) {

	tests := []struct {
		input         string
		expectedValue string
	}{
		{"let x = 3;", "x"},
		{"let y = 3;", "y"},
		// {"let foo_bar = bar;", "foo_bar"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		if len(p.errors) > 0 {
			for _, err := range p.errors {
				t.Errorf(err)
			}
			t.FailNow()
		}

		if len(stmts.Statements) != 1 {
			t.Fatalf("wrong number of stmts.Statements, expected 1, got %v instead", len(stmts.Statements))
		}
		letStmt, ok := stmts.Statements[0].(*ast.LetStatement)
		if !ok {
			t.Fatalf("stmts.Statements[0] is not *ast.LetStatement, got '%T' instead", stmts.Statements[0])
		}
		if letStmt.Name.Value != test.expectedValue {
			t.Fatalf("letStmt.Name is not '%s', got '%s' instead", test.expectedValue, letStmt.Name.Value)
		}
	}
}

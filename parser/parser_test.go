package parser

import (
	"testing"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
)

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

	ident, ok := exprStmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expected ExpressionStatement got %T instead", exprStmt.Expression)
	}

	if ident.Value != 5 {
		t.Fatalf("expected ident of value foo, got %v instead", ident.Value)
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

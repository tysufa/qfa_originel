package parser

import (
	"testing"

	"github.com/tysufa/qfa/ast"
	"github.com/tysufa/qfa/lexer"
)

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"
	l := lexer.New(input)
	p := New(l)
	program := p.GetStatements()
	testParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T",
			stmt.Expression)
	}

	if !testIdentifier(t, exp.Function, "add") {
		return
	}

	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}
	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.GetStatements()
		testParserErrors(t, p)
		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)
		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d\n",
				len(tt.expectedParams), len(function.Parameters))
		}
		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`
	l := lexer.New(input)
	p := New(l)
	program := p.GetStatements()
	testParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T",
			stmt.Expression)
	}
	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n",
			len(function.Parameters))
	}
	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")
	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n",
			len(function.Body.Statements))
	}
	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T",
			function.Body.Statements[0])
	}
	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

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
		{
			"a + add(b * c) + d",
			"((a+add((b*c)))+d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2*3), (4+5), add(6, (7*8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a+b)+((c*d)/f))+g))",
		},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		if len(p.Errors) > 0 {
			for _, err := range p.Errors {
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

		if len(p.Errors) > 0 {
			for _, err := range p.Errors {
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

		if len(p.Errors) > 0 {
			for _, err := range p.Errors {
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

func testInfixExpression(t *testing.T, infixExpr ast.Expression, left interface{}, operator string, right interface{}) {
	infix, ok := infixExpr.(*ast.InfixExpression)

	if !ok {
		t.Fatalf("expected *ast.InfixExpression got %T instead", infixExpr)
	}

	testLiteralExpression(t, infix.Left, left)

	if infix.Operator != operator {
		t.Fatalf("expected operator '%v', got '%v' instead", operator, infix.Operator)
	}

	testLiteralExpression(t, infix.Right, right)
}

func testBooleanLiteral(t *testing.T, booleanExpr ast.Expression, expectedBool bool) {
	bool, ok := booleanExpr.(*ast.Boolean)

	if !ok {
		t.Fatalf("expected *ast.Boolean got %T instead", booleanExpr)
	}
	if bool.Value != expectedBool {
		t.Fatalf("expected boolean of value '%v', got '%v' instead", expectedBool, bool.Value)
	}
}

func testIdentLiteral(t *testing.T, identExpr ast.Expression, expectedIdent string) {
	ident, ok := identExpr.(*ast.Identifier)

	if !ok {
		t.Fatalf("expected *ast.identifier got %T instead", identExpr)
	}
	if ident.Value != expectedIdent {
		t.Fatalf("expected ident of value %s, got %s instead", expectedIdent, ident.Value)
	}
}

func testIntegerLiteral(t *testing.T, integerExpression ast.Expression, expectedValue int) {
	integ, ok := integerExpression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expected IntegerLiteral got %T instead", integerExpression)
	}

	if integ.Value != expectedValue {
		t.Fatalf("wrong integer value, expected %v, got %v instead", expectedValue, integ.Value)
	}

}

func TestBooleanLiteralExpressions(t *testing.T) {
	input := "true;"

	l := lexer.New(input)
	p := New(l)

	stmt := p.GetStatements()

	testStatementsNumber(t, 1, stmt.Statements)

	exprStmt, ok := stmt.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("expected ExpressionStatement got %T instead", stmt.Statements[0])
	}

	bool, ok := exprStmt.Expression.(*ast.Boolean)

	if !ok {
		t.Fatalf("expected *ast.Boolean got %T instead", exprStmt.Expression)
	}

	if bool.Value != true {
		t.Fatalf("expected 'true', got '%v' instead", bool.Value)
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

	testIdentLiteral(t, exprStmt.Expression, "foo")

}

func TestWhileStatement(t *testing.T){
  tests := []struct{
    input string
    expected interface{}
  }{
    {"while (x < y){true}",true},
    {"while (x < y){x}", "x"},
  }

  for _, test := range tests{
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		testParserErrors(t, p)
		testStatementsNumber(t, 1, stmts.Statements)

    whileStmt, ok := stmts.Statements[0].(*ast.WhileStatement)

    if !ok{
      t.Fatalf("stmts.Satements[0] is not ast.WhileStatement, got %T instead", stmts.Statements[0])
    }

    testInfixExpression(t, whileStmt.Condition, "x", "<", "y")

    for i, stmt  := range whileStmt.Instructions.Statements{
      exprStmt, ok := stmt.(*ast.ExpressionStatement)
      if !ok{
        t.Fatalf("whileStmt.Instructions.Statements[%d] is not *ast.ExpressionStatement, got %T instead", i, stmt)
      }

      testLiteralExpression(t, exprStmt.Expression, test.expected)
    }
  }

}

func TestIfStatements(t *testing.T) {
	tests := []struct {
		input        string
		expected     interface{}
		expectedElse interface{}
	}{
		{"if (x < y){ true }", true, nil},
		{"if (x < y){ x }", "x", nil},
		{"if (x < y){ x } else { y }", "x", "y"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		testParserErrors(t, p)
		testStatementsNumber(t, 1, stmts.Statements)

		exprStmt, ok := stmts.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("expected ExpressionStatement got %T instead", stmts.Statements[0])
		}

		ifStmt, ok := exprStmt.Expression.(*ast.IfExpression)
		if !ok {
			t.Fatalf("stmts.Statements[0] is not *ast.IfExpression, got '%T' instead", stmts.Statements[0])
		}

		testInfixExpression(t, ifStmt.Condition, "x", "<", "y")

		for i, consequence := range ifStmt.Consequences.Statements {
			csq, ok := consequence.(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("whileStmt.Instructions.Statements[%d] is not *ast.ExpressionStatement, got %T instead", i, consequence)
			}
			testLiteralExpression(t, csq.Expression, test.expected)
		}

		if ifStmt.ElseConsequences != nil {
			for _, consequence := range ifStmt.ElseConsequences.Statements {
				csq, ok := consequence.(*ast.ExpressionStatement)
				if !ok {
					t.Fatalf("nop")
				}
				testLiteralExpression(t, csq.Expression, test.expectedElse)
			}
		}
	}
}

func TestReturnStatements(t *testing.T) {

	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return x;", "x"},
		{"return 1;", 1},
		{"return true;", true},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		testParserErrors(t, p)

		testStatementsNumber(t, 1, stmts.Statements)
		letStmt, ok := stmts.Statements[0].(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmts.Statements[0] is not *ast.ReturnStatement, got '%T' instead", stmts.Statements[0])
		}

		testLiteralExpression(t, letStmt.Value, test.expectedValue)

	}
}

func TestLetStatements(t *testing.T) {

	tests := []struct {
		input         string
		expectedName  string
		expectedValue interface{}
	}{
		{"let x = 3;", "x", 3},
		{"let z = false;", "z", false},
		{"let foo_bar = bar;", "foo_bar", "bar"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		stmts := p.GetStatements()

		testParserErrors(t, p)

		testStatementsNumber(t, 1, stmts.Statements)
		letStmt, ok := stmts.Statements[0].(*ast.LetStatement)
		if !ok {
			t.Fatalf("stmts.Statements[0] is not *ast.LetStatement, got '%T' instead", stmts.Statements[0])
		}
		if letStmt.Name.Value != test.expectedName {
			t.Fatalf("letStmt.Name is not '%s', got '%s' instead", test.expectedName, letStmt.Name.Value)
		}

		testLiteralExpression(t, letStmt.Value, test.expectedValue)

	}
}

func testLiteralExpression(t *testing.T, expression ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
		testIntegerLiteral(t, expression, v)
	case string:
		testIdentLiteral(t, expression, v)
	case bool:
		testBooleanLiteral(t, expression, v)
	}
}

func testParserErrors(t *testing.T, p *Parser) {
	if len(p.Errors) > 0 {
		for _, err := range p.Errors {
			t.Errorf("Parser error : " + err)
		}
		t.FailNow()
	}
}

func testStatementsNumber(t *testing.T, stmtNb int, stmts []ast.Statement) {
	if len(stmts) != stmtNb {
		t.Fatalf("wrong number of stmts.Statements, expected %d, got %v instead", stmtNb, len(stmts))
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

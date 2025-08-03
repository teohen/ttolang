package parser

import (
	"fmt"
	"testing"

	"github.com/teohen/ttolang/ast"
	"github.com/teohen/ttolang/lexer"
	"github.com/teohen/ttolang/utils"
)

func TestCriaStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"cria x <- 5;", "x", 5},
		{"cria y <- vdd;", "y", true},
		{"cria foobar <- y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]

		if !testCriaStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		criaStmt := stmt.(*ast.CriaStatement)
		if !testLiteralExpression(t, criaStmt.Value, tt.expectedValue) {
			return
		}
	}
}

func testCriaStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "cria" {
		t.Errorf("s.TokenLiteral not 'cria'. got=%q", s.TokenLiteral())
		return false
	}

	criaStmt, ok := s.(*ast.CriaStatement)

	if !ok {
		t.Errorf("s not *ast.CriaStatement. got=%T", s)
		return false
	}

	if criaStmt.Name.Value != name {
		t.Errorf("criaStmt.Name.Value not '%s', got=%s", name, criaStmt.Name.Value)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestDevolveStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"devolve 5;", 5},
		{"devolve true;", true},
		{"devolve foobar;", "foobar"},
	}

	for _, tt := range tests {

		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		devolveStmt, ok := stmt.(*ast.DevolveStatement)

		if !ok {
			t.Fatalf("sttm not *ast.DevolveStatement. got=%T", stmt)
		}

		if devolveStmt.TokenLiteral() != "devolve" {
			t.Fatalf("devolveStmt.TokenLiteral not 'devolve'. got=%q", devolveStmt.TokenLiteral())
		}

		if testLiteralExpression(t, devolveStmt.ReturnValue, tt.expectedValue) {
			return
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("expression  not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("expression  not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!vdd;", "!", true},
		{"!falso;", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 = 5;", 5, "=", 5},
		{"5 != 5;", 5, "!=", 5},
		{"vdd = vdd;", true, "=", true},
		{"vdd != falso;", true, "!=", false},
		{"falso = falso;", false, "=", false},
		{"vdd & vdd;", true, "&", true},
		{"vdd | vdd;", true, "|", true},
		{"vdd & falso;", true, "&", false},
		{"vdd | falso;", true, "|", false},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d. got=%d", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		if !testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 = 3 < 4",
			"((5 > 4) = (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 = 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) = ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 = 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) = ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 = falso",
			"((3 > 5) = falso)",
		},
		{
			"3 < 5 = vdd",
			"((3 < 5) = vdd)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true = true)",
			"(!(true = true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
		{
			"vdd & vdd & (vdd & vdd)",
			"((vdd & vdd) & (vdd & vdd))",
		},
		{
			"vdd & vdd & (falso | vdd)",
			"((vdd & vdd) & (falso | vdd))",
		},
		{
			"vdd & vdd = falso = (vdd & vdd)",
			"(vdd & ((vdd = falso) = (vdd & vdd)))",
		},
		{
			"vdd & vdd | vdd",
			"((vdd & vdd) | vdd)",
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"vdd;", true},
		{"falso;", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		boolean, ok := stmt.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("exp is not *ast.Boolean. got=%T", stmt.Expression)
		}

		if boolean.Value != tt.expectedBoolean {
			t.Errorf("boolean.Value not %t. got=%t", tt.expectedBoolean, boolean.Value)
		}
	}
}

func TestSeExpression(t *testing.T) {
	input := `se (x < y) { x }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.body does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%t", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.SeExpression)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.SeExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("exp.Consequence.Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
}

func TestSeSenaoExpression(t *testing.T) {
	input := `se (x < y) { x } senao { y }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.SeExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.SeExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n",
			len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if len(exp.Alternative.Statements) != 1 {
		t.Errorf("exp.Alternative.Statements does not contain 1 statements. got=%d\n",
			len(exp.Alternative.Statements))
	}

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			exp.Alternative.Statements[0])
	}

	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}
}

func TestProcLiteralParsing(t *testing.T) {
	input := `proc(x, y) {x + y;}`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.ProcLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.ProcLiteral. got=%T", stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n", len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T", function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "proc() {};", expectedParams: []string{}},
		{input: "proc(x) {};", expectedParams: []string{"x"}},
		{input: "proc(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.ProcLiteral)
		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d\n",
				len(tt.expectedParams), len(function.Parameters))
		}
		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
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

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)

	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %q. got=%q", "hello world", literal.Value)
	}
}

func TestParsingListaLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	lista, ok := stmt.Expression.(*ast.ListaLiteral)

	if !ok {
		t.Fatalf("exp not ast.ListaLiteral. got=%T", stmt.Expression)
	}

	if len(lista.Elements) != 3 {
		t.Fatalf("len(lista.Elements) not 3. got=%d", len(lista.Elements))
	}

	expectedInteger := int64(1)
	testIntegerLiteral(t, lista.Elements[0], expectedInteger)
	testInfixExpression(t, lista.Elements[1], 2, "*", 2)
	testInfixExpression(t, lista.Elements[2], 3, "+", 3)

}

func TestParsingIndexExpressions(t *testing.T) {
	input := "minha_lista[1 + 1];"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}

	if !testIdentifier(t, indexExp.Left, "minha_lista") {
		return
	}

	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}

}

func TestAssignStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
		expectedRight      interface{}
		expectedLeft       interface{}
	}{
		{"any <- 3 + 3;", "any", 6, 3, 3},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
		}

		stmt := program.Statements[0]

		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		assignSttm := stmt.(*ast.AssignStatement)

		if !testIdentifier(t, assignSttm.Name, "any") {
			return
		}

		if !testInfixExpression(t, assignSttm.AssignExpression, tt.expectedLeft, "+", tt.expectedRight) {
			return
		}

	}
}

func TestRepeteExpression(t *testing.T) {
	input := `repete(i <- i + 1 ate i < 10){
				i;
			  }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.body does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%t", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.RepeteExpression)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.RepeteExpression. got=%T", stmt.Expression)
	}

	if !testAssignStatement(t, exp.Step, "i") {
		return
	}

	if !testInfixExpression(t, exp.Condition, "i", "<", 10) {
		return
	}

	if len(exp.RepeatingStatements.Statements) != 1 {
		t.Errorf("RepeatingStatements is not 1 statements. got=%d\n", len(exp.RepeatingStatements.Statements))
	}

	repeatingStatement, ok := exp.RepeatingStatements.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("exp.RepeatingStatement.Statements[0] is not ast.ExpressionStatement. got=%T", exp.RepeatingStatements.Statements[0])
	}

	if !testIdentifier(t, repeatingStatement.Expression, "i") {
		return
	}
}

func TestParsingEstruturaLiterals(t *testing.T) {
	input := `
			{
			  nome <- "ttolang", 
			  cod <- 1,
			  op <- proc(x) { x; }
			};
			`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	estrutura, ok := stmt.Expression.(*ast.EstruturaLiteral)

	if !ok {
		t.Fatalf("exp not ast.EstruturaLiteral. got=%T", stmt.Expression)
	}

	if len(estrutura.Items) != 3 {
		t.Fatalf("len(estrutura.Items) not 3. got=%d", len(estrutura.Items))
	}

	expectedKeys := []string{"nome", "cod", "op"}

	for _, expk := range expectedKeys {
		_, ok := estrutura.Items[expk]
		if !ok {
			t.Fatalf("Estrutura.Items has wrong key. Expected=%s", expk)
		}
	}

	if !testStringLiteral(t, estrutura.Items["nome"], "ttolang") {
		return
	}

	if !testIntegerLiteral(t, estrutura.Items["cod"], int64(1)) {
		return
	}

	procLiteral, ok := estrutura.Items["op"].(*ast.ProcLiteral)
	if !ok {
		t.Fatalf("Estrutura.Items[op] is not a *ast.ProcLiteral. Got=%T", estrutura.Items["op"])
	}

	if len(procLiteral.Parameters) != 1 {
		t.Fatalf("procLiteral parameters wrong. want 1, got=%d", len(procLiteral.Parameters))
	}

	testLiteralExpression(t, procLiteral.Parameters[0], "x")

	if len(procLiteral.Body.Statements) != 1 {
		t.Fatalf("procLiteral.Body.Statements has not 1 statements. got=%d", len(procLiteral.Body.Statements))
	}

	bodyStmt, ok := procLiteral.Body.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("procLiteral body stmt is not ast.ExpressionStatement. got=%T", procLiteral.Body.Statements[0])
	}

	if !testLiteralExpression(t, bodyStmt.Expression, "x") {
		return
	}
}

func TestParsingEstruturaIndex(t *testing.T) {
	input := `
				{op <- proc(x) { x + 2; }}["op"](2)
			`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)

	if !ok {
		t.Fatalf("exp not ast.CallExpresssion. got=%T", stmt.Expression)
	}

	indexExpression, ok := exp.Function.(*ast.IndexExpression)

	if !ok {
		t.Fatalf("exp.Function not a *ast.IndexExpression. Got=%T", exp.Function)
	}

	if !testStringLiteral(t, indexExpression.Index, "op") {
		return
	}

	estrutura, ok := indexExpression.Left.(*ast.EstruturaLiteral)

	if !ok {
		t.Fatalf("indexExpression.Left not *ast.EstruturaLiteral. got=%T", stmt.Expression)
	}

	if len(estrutura.Items) != 1 {
		t.Fatalf("len(estrutura.Items) not 1. got=%d", len(estrutura.Items))
	}

	expectedKeys := []string{"op"}

	for _, expk := range expectedKeys {
		_, ok := estrutura.Items[expk]
		if !ok {
			t.Fatalf("Estrutura.Items has wrong key. Expected=%s", expk)
		}
	}

	procLiteral, ok := estrutura.Items["op"].(*ast.ProcLiteral)
	if !ok {
		t.Fatalf("Estrutura.Items[op] is not a *ast.ProcLiteral. Got=%T", estrutura.Items["op"])
	}

	if len(procLiteral.Parameters) != 1 {
		t.Fatalf("procLiteral parameters wrong. want 1, got=%d", len(procLiteral.Parameters))
	}

	testLiteralExpression(t, procLiteral.Parameters[0], "x")

	if len(procLiteral.Body.Statements) != 1 {
		t.Fatalf("procLiteral.Body.Statements has not 1 statements. got=%d", len(procLiteral.Body.Statements))
	}

	bodyStmt, ok := procLiteral.Body.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("procLiteral body stmt is not ast.ExpressionStatement. got=%T", procLiteral.Body.Statements[0])
	}

	if !testInfixExpression(t, bodyStmt.Expression, "x", "+", 2) {
		return
	}

}

func TestImportaStatement(t *testing.T) {
	packageCode := `
		cria nome <- "teteo";
	`

	err := utils.WriteFile("./biblioteca.tto", packageCode)

	if err != nil {
		t.Fatalf("Fail to create package file. got=%s", err.Error())
	}

	input := `
				importa "./biblioteca.tto";
				cria idade <- 3;
				mostra(nome);
			`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	packageSttms := program.Statements[:1]
	currentSttms := program.Statements[1:]

	if len(packageSttms) != 1 {
		t.Fatalf("package length not 1. got=%d", len(packageSttms))
	}

	if len(currentSttms) != 3 {
		fmt.Println(currentSttms)
		t.Fatalf("currenSttms length not 3. got=%d", len(currentSttms))
	}

	stmt, ok := currentSttms[0].(*ast.ImportaStatement)

	if !ok {
		t.Fatalf("currentStatements[0] not ast.ImportaStatement. got=%T", currentSttms[0])
	}

	if !testStringLiteral(t, stmt.FilePath, "./biblioteca.tto") {
		t.Fatalf("stmt.FilePath not a ast.StringLiteral. got=%T", stmt.FilePath)
	}

	if len(stmt.Program.Statements) != 2 {
		t.Fatalf("stmt.Program.Statement length not 2. got=%d", len(stmt.Program.Statements))
	}

	criaSttm, ok := stmt.Program.Statements[0].(*ast.CriaStatement)

	if !ok {
		t.Fatalf("criaSttm not ast.CriaStatement type. got=%T", criaSttm)
	}

	if !testCriaStatement(t, criaSttm, "nome") {
		return
	}

	if !testStringLiteral(t, criaSttm.Value, "teteo") {
		return
	}

}

func testStringLiteral(t *testing.T, sl ast.Expression, value interface{}) bool {

	stringLiteral, ok := sl.(*ast.StringLiteral)

	if !ok {
		t.Errorf("s not *ast.StringLiteral. got=%T", sl)
		return false
	}

	if stringLiteral.Value != value {
		t.Errorf("stringLiteral.Value not '%s', got=%s", value, stringLiteral.Value)
		return false
	}
	return true
}

func testAssignStatement(t *testing.T, s ast.Statement, name string) bool {
	assignSttm, ok := s.(*ast.AssignStatement)

	if !ok {
		t.Errorf("s not *ast.AssignStatement. got=%T", s)
		return false
	}

	if assignSttm.Name.Value != name {
		t.Errorf("assignSttm.Name.Value not '%s', got=%s", name, assignSttm.Name.Value)
		return false
	}
	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value interface{}) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("il is not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral  not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
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
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)

	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}

	literal := ""
	if value == true {
		literal = "vdd"
	} else {
		literal = "falso"
	}

	if bo.TokenLiteral() != literal {
		t.Errorf("bo.TokenLiteral not %t. got=%s", value, bo.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

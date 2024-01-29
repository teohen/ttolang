package evaluator

import (
	"testing"

	"github.com/teohen/ttolang/lexer"
	"github.com/teohen/ttolang/object"
	"github.com/teohen/ttolang/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
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
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{"vdd", true},
		{"mentira", false},

		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"vdd == vdd", true},
		{"mentira == mentira", true},
		{"vdd == mentira", false},
		{"vdd != mentira", true},
		{"mentira != vdd", true},
		{"(1 < 2) == vdd", true},
		{"(1 < 2) == mentira", false},
		{"(1 > 2) == vdd", false},
		{"(1 > 2) == mentira", true},
	}

	for _, tt := range tests {
		evaluted := testEval(tt.input)
		testBooleanObject(t, evaluted, tt.expect)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!vdd", false},
		{"!mentira", true},
		{"!5", false},
		{"!!vdd", true},
		{"!!mentira", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluted := testEval(tt.input)
		testBooleanObject(t, evaluted, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"se (vdd) { 10 }", 10},
		{"se (mentira) { 10 }", nil},
		{"se (1) { 10 }", 10},
		{"se (1 < 2) { 10 }", 10},
		{"se (1 > 2) { 10 }", nil},
		{"se (1 > 2) { 10 } senao { 20 }", 20},
		{"se (1 < 2) { 10 } senao { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		integer, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}

}

func TestDevolveStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"devolve 10;", 10},
		{"devolve 10; 9;", 10},
		{"devolve 2 * 5; 9;", 10},
		{"9; devolve 2 * 5; 9;", 10},
		{`se (10 > 1) {
			se (10 > 1) {
				devolve 10;
			}
			devolve 1;
		  }
		`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + vdd;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + vdd; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-vdd",
			"unknown operator: -BOOLEAN",
		},
		{
			"vdd + mentira;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; vdd + mentira; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { vdd + mentira; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{`
			se (10 > 1) {
				se (10 > 1) {
					devolve vdd + mentira;
				}
				devolve 1;
			}`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{"foobar", "identifier not found: foobar"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)

		if !ok {
			t.Errorf("no error object returned. got=%T (%v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}

}

func TestCriaStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"cria a = 5; a;", 5},
		{"cria a = 5 * 5; a;", 25},
		{"cria a = 5; cria b = a; b;", 5},
		{"cria a = 5; cria b = a; cria c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%v)", obj, obj)
		return false
	}
	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	restul, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not integer. got=%d, want=%d", restul.Value, expected)
		return false
	}

	if restul.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", restul.Value, expected)
		return false
	}

	return true
}
func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object was wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}

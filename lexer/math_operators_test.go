package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestMathOperators(t *testing.T) {
	input := `
	!-/*5;
	5 < 10 > 5;
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.typeExpected {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.typeExpected, tok.Literal)
		}
	}
}

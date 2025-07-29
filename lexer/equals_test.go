package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestEquals(t *testing.T) {
	input := `
	10 = 10;
	10 != 9;
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.INT, "10"},
		{token.EQ, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
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

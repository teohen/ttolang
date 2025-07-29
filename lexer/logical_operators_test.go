package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	vdd & vdd;
	falso | vdd;
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.VDD, "vdd"},
		{token.AND, "&"},
		{token.VDD, "vdd"},
		{token.SEMICOLON, ";"},
		{token.FALSE, "falso"},
		{token.OR, "|"},
		{token.VDD, "vdd"},
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

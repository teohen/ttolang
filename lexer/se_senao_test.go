package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestSeSenao(t *testing.T) {
	input := `
	se (5 < 10) {
		devolve vdd;
	} senao {
		devolve falso;
	}
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.IF, "se"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "devolve"},
		{token.VDD, "vdd"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "senao"},
		{token.LBRACE, "{"},
		{token.RETURN, "devolve"},
		{token.FALSE, "falso"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
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

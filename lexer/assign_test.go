package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestAssign(t *testing.T) {
	input := `
	cria i <- 1;
	i <- i + 1;
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.LET, "cria"},
		{token.IDENT, "i"},
		{token.ASSIGN, "<-"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.ASSIGN, "<-"},
		{token.IDENT, "i"},
		{token.PLUS, "+"},
		{token.INT, "1"},
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

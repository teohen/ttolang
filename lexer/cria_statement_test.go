package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestCriaStatement(t *testing.T) {
	input := `
	cria five <- 5;
	cria ten <- 10;
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.LET, "cria"},
		{token.IDENT, "five"},
		{token.ASSIGN, "<-"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "cria"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "<-"},
		{token.INT, "10"},
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

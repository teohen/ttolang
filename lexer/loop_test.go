package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestLoop(t *testing.T) {
	input := `
	repete(i <- i + 1 ate i < 10) {
		i;
	}
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.LOOP, "repete"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.ASSIGN, "<-"},
		{token.IDENT, "i"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.TO, "ate"},
		{token.IDENT, "i"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "i"},
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

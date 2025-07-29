package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestLista(t *testing.T) {
	input := `
	[1, 2];
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
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

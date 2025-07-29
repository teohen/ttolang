package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestStringLiteral(t *testing.T) {
	input := `
	"foobar"
	"foo bar"
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
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

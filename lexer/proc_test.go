package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestProcs(t *testing.T) {
	input := `
	cria five <- 5;
	cria ten <- 10;

	cria add <- proc(x,y) {
		x + y
	};

	cria result <- add(five, ten);
	!-/*5;
	5 < 10 > 5;

	se (5 < 10) {
		devolve vdd;
	} senao {
		devolve falso;
	}

	10 = 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];

	cria i <- 1;
	i <- i + 1;

	repete(i <- i + 1 ate i < 10) {
		i;
	}

	{nome <- "ttolang", code <- 1, op <- proc(x) { x; }}
	{op <- proc(x) { x + 2; }}["op"](2);
	vdd & vdd;
	falso | vdd;
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
		{token.LET, "cria"},
		{token.IDENT, "add"},
		{token.ASSIGN, "<-"},
		{token.FUNCTION, "proc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "cria"},
		{token.IDENT, "result"},
		{token.ASSIGN, "<-"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
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
		{token.INT, "10"},
		{token.EQ, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
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
		{token.LBRACE, "{"},
		{token.IDENT, "nome"},
		{token.ASSIGN, "<-"},
		{token.STRING, "ttolang"},
		{token.COMMA, ","},
		{token.IDENT, "code"},
		{token.ASSIGN, "<-"},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.IDENT, "op"},
		{token.ASSIGN, "<-"},
		{token.FUNCTION, "proc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.LBRACE, "{"},
		{token.IDENT, "op"},
		{token.ASSIGN, "<-"},
		{token.FUNCTION, "proc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.STRING, "op"},
		{token.RBRACKET, "]"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
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

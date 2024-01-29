package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = ":>"
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	EQ        = "="
	NEQ       = "!="

	FUNCTION = "FUNCTION"
	LET      = "LET"
	VDD      = "VDD"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	LT = "<"
	GT = ">"

	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"proc":    FUNCTION,
	"cria":    LET,
	"vdd":     VDD,
	"mentira": FALSE,
	"se":      IF,
	"senao":   ELSE,
	"devolve": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

package token

type TokenType string

const (
	BANG      = "!"
	PLUS      = "+"
	MINUS     = "-"
	SLASH     = "/"
	STAR      = "*"
	LPAR      = "("
	RPAR      = ")"
	LBR       = "{"
	RBR       = "}"
	EQ        = "="
	IDENT     = "IDENT"
	INT       = "INT"
	FLOAT     = "FLOAT"
	LET       = "LET"
	SEMICOLON = ";"
	EOF       = "EOF"
	NL        = "NL" // New Line
)

type Token struct {
	Value string
	Type  TokenType
	Line  int
}

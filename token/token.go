package token

type TokenType string

const (
	BANG  = "!"
	PLUS  = "+"
	MINUS = "-"
	SLASH = "/"
	STAR  = "*"
	LPAR  = "("
	RPAR  = ")"
	LBR   = "{"
	RBR   = "}"
	EQ    = "="
)

type Token struct {
	Value string
	Type  TokenType
	Line  int
}

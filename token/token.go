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
	GT        = ">"
	LT        = "<"
	GEQT      = ">="
	LEQT      = "<="
	EQEQ      = "=="
	NEQ       = "!="
	IDENT     = "IDENT"
	INT       = "INT"
	FLOAT     = "FLOAT"
	IF        = "IF"
	ELSE      = "ELSE"
	FN        = "FN"
	WHILE     = "WHILE"
	PRINT     = "PRINT"
	RETURN    = "RETURN"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	LET       = "LET"
	SEMICOLON = ";"
	COMMA     = ","
	EOF       = "EOF"
	NL        = "NL" // New Line
	ILLEGAL   = "ILLEGAL"
)

var Reserved = map[string]TokenType{
	"fn":     FN,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"let":    LET,
	"while":  WHILE,
	"print":  PRINT,
}

type Token struct {
	Value string
	Type  TokenType
	Line  int
}

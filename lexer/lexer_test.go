package lexer

import (
	"testing"

	"github.com/tysufa/qfa/token"
)

func TestGetToken(t *testing.T) {
	input := `!+/*-(){}`

	l := New(input)

	tests := []struct {
		expectedValue string
		expectedType  token.TokenType
	}{
		{"!", token.BANG},
		{"+", token.PLUS},
		{"/", token.SLASH},
		{"*", token.STAR},
		{"-", token.MINUS},
		{"(", token.LPAR},
		{")", token.RPAR},
		{"{", token.LBR},
		{"}", token.RBR},
	}

	for _, tt := range tests {
		tok := l.getToken()
		if tt.expectedType != tok.Type {
			t.Fatalf("wrong token type, expected %s, got %s instead", tt.expectedType, tok.Type)
		}
		if tt.expectedValue != tok.Value {
			t.Fatalf("wrong token value, expected %s, got %s instead", tt.expectedValue, tok.Value)
		}
	}

}

func TestNextChar(t *testing.T) {
	input := `foo bar`

	test := [7]byte{'f', 'o', 'o', ' ', 'b', 'a', 'r'}

	l := New(input)

	for _, char := range test {
		if char != l.curChar {
			t.Fatalf("wrong character, expected '%c', got '%c' instead.", char, l.curChar)
		}
		l.nextChar()
	}
}

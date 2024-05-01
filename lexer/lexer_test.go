package lexer

import (
	"testing"

	"github.com/tysufa/qfa/token"
)

func TestGetToken(t *testing.T) {
	input := `let foo = bar;`

	l := New(input)

	tests := []struct {
		expectedValue string
		expectedType  token.TokenType
	}{
		{"let", token.LET},
		{"foo", token.IDENT},
		{"=", token.EQ},
		{"bar", token.IDENT},
		{";", token.SEMICOLON},
		{"", token.EOF},
	}

	for _, tt := range tests {
		tok := l.getToken()
		if tt.expectedType != tok.Type {
			t.Fatalf("wrong token type, expected '%s', got '%s' instead", tt.expectedType, tok.Type)
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

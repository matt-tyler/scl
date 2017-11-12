package lexer

import (
	"strings"
	"testing"

	"github.com/matt-tyler/scl/pkg/token"
)

func TestScanner(t *testing.T) {

	var tests = []struct {
		description string
		input       string
		expected    []token.TokenType
	}{
		{"empty case", "", []token.TokenType{}},
		{"single case", "doug", []token.TokenType{token.String}},
		{"single case w/ whitespace", "  doug ", []token.TokenType{token.String}},
		{"multi case", "doug on monday, tuesday", []token.TokenType{
			token.String, token.On, token.String, token.Comma, token.String}},
		{"parenthesis", "doug on (monday and tuesday)", []token.TokenType{
			token.String, token.On, token.LeftParen, token.String, token.And, token.String, token.RightParen}},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		lexer := New(input)

		for _, expected := range tt.expected {

			if !lexer.Next() {
				t.Errorf("%v: A token is missing from the output", tt.description)
			}

			eTok := &token.Token{Typ: expected}

			if tok := lexer.Token(); tok.Typ != expected {
				t.Errorf("%v: Expected %v, actually %v", tt.description, eTok.String(), tok.String())
			}
		}
	}
}

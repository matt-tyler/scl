package parser

import (
	"strings"
	"testing"

	"github.com/matt-tyler/scl/pkg/lexer"
	"github.com/matt-tyler/scl/pkg/token"
)

func TestPCParser(t *testing.T) {
	var table = []struct {
		description string
		input       string
		err         bool
		pos         int
	}{
		{"successful case", "Between 2006-01-02T15:04:05Z 2007-01-02T15:04:05Z", false, 3},
		{"mixed on & between", "On wednesdays, weekends and Between 2006-01-02T15:04:05Z 2007-01-02T15:04:05Z", false, 8},
		//{"parenthesis", "On mondays or (On Tuesdays and Between 2006-01-02T15:04:05Z 2007-01-02T15:04:05Z)", false, 11},
		{"mismatched parenthesis", "On mondays or (On Tuesdays and Between 2006-01-02T15:04:05Z 2007-01-02T15:04:05Z))", true, 11},
		{"missing parenthesis", "On mondays or (On Tuesdays and Between 2006-01-02T15:04:05Z 2007-01-02T15:04:05Z", true, 10},
		{"bad token", "Wom", true, 0},
		{"bad token 2", "On Wom", true, 1},
		{"bad token 3", "On Mondays, and tuesdays", true, 3},
	}
	for _, tt := range table {
		lexer := lexer.New(strings.NewReader(tt.input))
		tokens := []token.Token{}

		for lexer.Next() {
			tokens = append(tokens, lexer.Token())
		}

		p := &pcParser{0, tokens}
		_, pos, err := p.ParseExpr()

		if pos != tt.pos {
			t.Errorf("case %v: Expected to finish parsing at position %v, instead %v", tt.description, tt.pos, pos)
		}

		if tt.err {
			if err == nil {
				t.Errorf("case %v: An error should have occurred, but did not", tt.description)
				continue
			}
		} else {
			if err != nil {
				t.Errorf("An unexpected error occurred: %v", err)
				continue
			}
		}
	}
}

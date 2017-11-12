package parser

import (
	"strings"
	"testing"
	"time"

	"github.com/matt-tyler/scl/pkg/lexer"
	"github.com/matt-tyler/scl/pkg/token"
)

func TestBetweenEval(t *testing.T) {
	var table = []struct {
		description string
		from        time.Time
		until       time.Time
		param       time.Time
		result      bool
	}{
		{"In between", time.Now(), time.Now().Add(time.Hour), time.Now().Add(time.Minute), true},
		{"After", time.Now(), time.Now().Add(time.Minute), time.Now().Add(time.Hour), false},
		{"Before", time.Now(), time.Now().Add(time.Minute), time.Now().Add(-5 * time.Hour), false},
	}

	for _, tt := range table {

		b := Between{tt.from, tt.until}
		if b.Eval(tt.param) != tt.result {
			t.Errorf("Case %v failed", tt.description)
		}
	}
}

func TestBetween(t *testing.T) {
	var table = []struct {
		description string
		input       string
		err         bool
		expected    []string
	}{
		{"successful case", "2006-01-02T15:04:05Z 2007-01-02T15:04:05Z", false, nil},
		{"missing date", "2006-01-02T15:04:05Z", true, nil},
		{"empty string", "", true, nil},
		{"bad input", "asdkasd alsdkalskd", true, nil},
	}
	for _, tt := range table {
		lexer := lexer.New(strings.NewReader(tt.input))
		tokens := []token.Token{}

		for lexer.Next() {
			tokens = append(tokens, lexer.Token())
		}

		p := &pcParser{0, tokens}

		b := Between{}
		if err := ParseBetween(p, &b); tt.err {
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

package parser

import (
	"strings"
	"testing"
	"time"

	"github.com/matt-tyler/scl/pkg/lexer"
	"github.com/matt-tyler/scl/pkg/token"
)

func TestOnEval(t *testing.T) {

	monday := time.Date(2017, 11, 13, 12, 0, 0, 0, time.Local)
	tuesday := monday.Add(24 * time.Hour)
	wednesday := tuesday.Add(24 * time.Hour)
	thursday := wednesday.Add(24 * time.Hour)
	friday := thursday.Add(24 * time.Hour)
	saturday := friday.Add(24 * time.Hour)
	sunday := saturday.Add(24 * time.Hour)

	var table = []struct {
		description string
		periods     []Period
		param       time.Time
		result      bool
	}{
		{"weekday", []Period{7}, tuesday, true},
		{"not weekday", []Period{7}, sunday, false},
		{"weekend", []Period{8}, sunday, true},
		{"not weekend", []Period{8}, wednesday, false},
		{"specific day", []Period{2}, tuesday, true},
		{"not specific day", []Period{2}, wednesday, false},
	}

	for _, tt := range table {
		on := On{tt.periods}
		if on.Eval(tt.param) != tt.result {
			t.Errorf("Case %v failed", tt.description)
		}
	}
}

func TestParseOn(t *testing.T) {
	var table = []struct {
		description string
		input       string
		err         bool
		expected    []string
	}{
		{"singular case", "mondays", false, []string{"mondays"}},
		{"trailing comma", "mondays,", true, nil},
		{"multiple literals", "mondays, tuesdays", false, []string{"mondays", "tuesdays"}},
		{"multiple literals /w trailing comma", "mondays, tuesdays,", true, nil},
		{"invalid literal", "pancakes", true, nil},
		{"invalid literal multi", "mondays, pancakes", true, nil},
		{"", "mondays and", false, []string{"mondays"}},
	}

	for _, tt := range table {
		lexer := lexer.New(strings.NewReader(tt.input))
		tokens := []token.Token{}

		for lexer.Next() {
			tokens = append(tokens, lexer.Token())
		}

		p := &pcParser{0, tokens}

		o := On{}
		if err := ParseOn(p, &o); tt.err {
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
		if len(o.periods) != len(tt.expected) {
			t.Errorf("Case %v: expected %v, actually %v", tt.description, tt.expected, o.periods)
			continue
		}

		for i, v := range o.periods {
			tok := token.Token{Typ: token.String, Val: tt.expected[i]}
			var period Period
			err := TokenToPeriod(tok, &period)
			if err != nil {
				t.Errorf("Unexpected error occurred: %v", err)
			}
			if v != period {
				t.Errorf("Case %v: expected %v, actually %v", tt.description, tt.expected, o.periods)
			}
		}
	}
}

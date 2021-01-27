package evaluator

import (
	"testing"

	"github.com/AlexKondov/go-task/internal/token"
)

func TestTokenEvaluation(t *testing.T) {
	tests := []struct {
		name   string
		input  []token.Token
		expect int
	}{
		{"No operation", []token.Token{
			{"number", "5"},
		}, 5},
		{"Single operation", []token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
		}, 9},
		{"Two operations", []token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
			{"operator", token.DIVIDE},
			{"number", "3"},
		}, 3},
		{"Multiple operations", []token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
			{"operator", token.DIVIDE},
			{"number", "3"},
			{"operator", token.MULTIPLY},
			{"number", "2"},
		}, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			e := New(test.input)
			n := e.Evaluate()

			if n != test.expect {
				t.Fatalf("evaluator.Evaluate() produced wrong result: %d", n)
			}
		})
	}
}

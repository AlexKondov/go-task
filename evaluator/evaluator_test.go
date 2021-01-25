package evaluator

import (
	"go-task/token"
	"testing"
)

func TestTokenEvaluation(t *testing.T) {
	tests := []struct {
		input  []token.Token
		expect int
	}{
		{[]token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
		}, 9},
		{[]token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
			{"operator", token.DIVIDE},
			{"number", "3"},
		}, 3},
		{[]token.Token{
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
		e := New(test.input)
		n := e.Evaluate()

		if n != test.expect {
			t.Fatalf("evaluator.Evaluate() produced wrong result: %d", n)
		}
	}
}

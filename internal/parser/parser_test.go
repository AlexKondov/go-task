package parser

import (
	"testing"

	"github.com/AlexKondov/go-task/internal/token"
)

func TestExpressionParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{"What is 5?", []token.Token{
			{"number", "5"},
		}},
		{"What is 5 plus 4?", []token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
		}},
		{"What is 5 plus 4 divided by 3?", []token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
			{"operator", token.DIVIDE},
			{"number", "3"},
		}},
		{"What is 5 plus 4 divided by 3 multiplied by 2?", []token.Token{
			{"number", "5"},
			{"operator", token.PLUS},
			{"number", "4"},
			{"operator", token.DIVIDE},
			{"number", "3"},
			{"operator", token.MULTIPLY},
			{"number", "2"},
		}},
	}

	for _, test := range tests {
		p := New(test.input)
		tokens, _ := p.ParseExpression()

		for i, token := range tokens {
			if token != test.expected[i] {
				t.Fatalf("parser.ParseExpression() produced wrong tokens: %s", token)
			}
		}
	}
}

func TestExpressionParsingErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"What is 5 plus 4", "evaluation not terminated properly"},
		{"What is 5 plus 4 divided by?", "evaluation should not end on an operator"},
		{"What is 5 plus plus 4?", "Keyword plus plus is not supported"},
		{"What is plus 4?", "expression should start with a number"},
	}

	for _, test := range tests {
		p := New(test.input)
		_, err := p.ParseExpression()

		if err.Error() != test.expected {
			t.Fatalf("parser.ParseExpression() produced wrong error: %s", err.Error())
		}
	}
}

package evaluator

import (
	"go-task/token"
	"strconv"
)

type Evaluator struct {
	tokens []token.Token
}

func New(tokens []token.Token) *Evaluator {
	e := &Evaluator{
		tokens: tokens,
	}

	return e
}

func (e *Evaluator) Evaluate() int {
	result, _ := strconv.Atoi(e.tokens[0].Value)
	var currentOperator string

	for _, t := range e.tokens {
		if t.Type == token.OPERATOR {
			currentOperator = t.Value
			continue
		}

		if t.Type == token.NUMBER && currentOperator != "" {
			// Execute current operator with the two numbers
			n, _ := strconv.Atoi(t.Value)
			result = calculate(currentOperator, result, n)
			currentOperator = ""
		}
	}

	return result
}

func calculate(operator string, left, right int) int {
	var n int

	switch operator {
	case token.PLUS:
		n = left + right
		break
	case token.MINUS:
		n = left - right
		break
	case token.DIVIDE:
		n = left / right
		break
	case token.MULTIPLY:
		n = left * right
		break
	}

	return n
}

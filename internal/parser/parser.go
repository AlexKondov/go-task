package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/AlexKondov/go-task/internal/token"
)

type Parser struct {
	expression string
}

func New(expression string) *Parser {
	p := &Parser{
		expression: expression,
	}

	return p
}

func (p *Parser) ParseExpression() ([]token.Token, error) {
	e := strings.ToLower(p.expression)

	if !strings.HasPrefix(e, token.QUESTION) {
		return nil, errors.New("evaluation must start with: " + token.QUESTION)
	}

	if !strings.HasSuffix(e, token.TERMINATION_SIGN) {
		return nil, errors.New("evaluation not terminated properly")
	}

	text := e[len(token.QUESTION) : len(e)-1]
	words := strings.Fields(text)

	var tokens []token.Token
	var operator string

	for i := 0; i < len(words); i++ {
		word := words[i]

		_, err := strconv.Atoi(word)

		if err != nil {
			// We append the string to the operator until we hit another number
			operator += " " + word
			continue
		}

		if operator != "" {
			if v, ok := token.Keywords[strings.TrimSpace(operator)]; ok {
				t := token.NewOperatorToken(v)
				tokens = append(tokens, t)
				operator = ""
			} else {
				// It's not a supported value or keyword
				return nil, errors.New("Keyword" + operator + " is not supported")
			}
		}

		// The token is a number
		t := token.NewNumericToken(word)
		tokens = append(tokens, t)
	}

	if operator != "" {
		// If an unhandled operator has remained, add it to the stack
		if v, ok := token.Keywords[strings.TrimSpace(operator)]; ok {
			t := token.NewOperatorToken(v)
			tokens = append(tokens, t)
			operator = ""
		} else {
			// It's not a supported value or keyword
			return nil, errors.New("Keyword" + operator + " is not supported")
		}
	}

	// Check if the last element is a number, if it's not it's not a valid operation
	// We do this after the parsing to make sure we show errors in the order in which they occur
	if tokens[len(tokens)-1].Type != token.NUMBER {
		return nil, errors.New("evaluation should not end on an operator")
	}

	return tokens, nil
}

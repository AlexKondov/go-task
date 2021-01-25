package api

import (
	"go-task/evaluator"
	"go-task/parser"
	"go-task/storage"
)

func IsValidExpression(expression string) (bool, error) {
	p := parser.New(expression)
	_, err := p.ParseExpression()

	if err != nil {
		return false, err
	}

	return true, nil
}

func EvaluateExpression(expression string) (int, error) {
	p := parser.New(expression)
	tokens, err := p.ParseExpression()

	if err != nil {
		return 0, err
	}

	e := evaluator.New(tokens)
	result := e.Evaluate()

	return result, nil
}

func GetExpressionErrors() []storage.ExpressionError {
	return storage.ErrorStorage.GetErrors()
}

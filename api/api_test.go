package api

import (
	"github.com/AlexKondov/go-task/internal/storage"
	"testing"
)

func TestEvaluateExpression(t *testing.T) {
	expression := "What is 5 plus 4?"

	num, err := EvaluateExpression(expression)

	if err != nil {
		t.Fatal("api.EvaluateExpression() failed")
	}

	if num != 9 {
		t.Fatalf("api.EvaluateExpression returned %d, expected 9", num)
	}
}

func TestEvaluateExpression_Error(t *testing.T) {
	expression := "What is 5 plus 4"

	num, err := EvaluateExpression(expression)

	if err == nil {
		t.Fatal("api.EvaluateExpression() error handling failed")
	}

	if num != 0 {
		t.Fatalf("api.EvaluateExpression returned %d, expected 0", num)
	}
}

func TestIsValidExpression(t *testing.T) {
	expression := "What is 5 plus 4?"

	valid, err := IsValidExpression(expression)

	if err != nil {
		t.Fatal("api.IsValidExpression() failed")
	}

	if valid != true {
		t.Fatalf("api.IsValidExpression returned false, expected true")
	}
}

func TestIsValidExpression_Error(t *testing.T) {
	expression := "What is 5 plus 4"

	valid, err := IsValidExpression(expression)

	if err == nil {
		t.Fatal("api.IsValidExpression() failed")
	}

	if valid == true {
		t.Fatalf("api.IsValidExpression returned true, expected false")
	}
}

func TestGetExpressionErrors(t *testing.T) {
	s := storage.New()
	s.SaveError("What is plus 4?", "/evaluate", "Invalid")

	errors := GetExpressionErrors(s)

	if len(errors) != 1 {
		t.Fatalf("api.GetExpressionErrors returned wrong number of errors")
	}
}
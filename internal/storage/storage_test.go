package storage

import "testing"

func TestErrorStorage(t *testing.T) {
	InitStorage()

	ErrorStorage.SaveError("What is 5 plus 4", "/test", "Not terminated properly")

	errs := ErrorStorage.GetErrors()

	if errs[0].Expression != "What is 5 plus 4" {
		t.Fatalf("storage.ErrorStorage.SaveError() saved wrong error: %s", errs[0].Expression)
	}

	ErrorStorage.SaveError("What is 5 plus plus 4", "/test", "Operator plus plus not supported")

	errs = ErrorStorage.GetErrors()

	if len(errs) != 2 {
		t.Fatalf("storage.ErrorStorage.GetErrors() returned wrong size: %d", len(errs))
	}

	ErrorStorage.SaveError("What is 5 plus 4", "/test", "Not terminated properly")

	if errs[0].Frequency != 2 {
		t.Fatalf("storage.ErrorStorage.SaveError() did not update error frequency: %s", errs[0].Expression)
	}
}

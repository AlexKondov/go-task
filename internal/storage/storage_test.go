package storage

import "testing"

func TestErrorStorage(t *testing.T) {
	s := New()
	s.SaveError("What is 5 plus 4", "/test", "Not terminated properly")

	errs := s.GetErrors()

	if errs[0].Expression != "What is 5 plus 4" {
		t.Fatalf("storage.ErrorStorage.SaveError() saved wrong error: %s", errs[0].Expression)
	}
}

func TestMultipleErrorStorage(t *testing.T) {
	s := New()

	s.SaveError("What is 5 plus 4", "/test", "Not terminated properly")
	s.SaveError("What is 5 plus plus 4", "/test", "Operator plus plus not supported")
	s.SaveError("What is 5 plus plus", "/test", "Not terminated properly")

	errs := s.GetErrors()

	if len(errs) != 3 {
		t.Fatalf("storage.ErrorStorage.GetErrors() returned wrong size: %d", len(errs))
	}
}

func TestErrorFrequency(t *testing.T) {
	s := New()

	s.SaveError("What is 5 plus 4", "/test", "Not terminated properly")
	s.SaveError("What is 5 plus 4", "/test", "Not terminated properly")
	s.SaveError("What is 5 plus plus 4", "/test", "Operator plus plus not supported")

	errs := s.GetErrors()

	if errs[0].Frequency != 2 {
		t.Fatalf("storage.ErrorStorage.SaveError() did not update error frequency: %s", errs[0].Expression)
	}
}

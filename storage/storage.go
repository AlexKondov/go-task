package storage

import "sync"

type ExpressionError struct {
	Expression string
	Endpoint   string
	Frequency  int
	Type       string
}

type Storage struct {
	errors []ExpressionError
	mutex  sync.Mutex
}

var ErrorStorage *Storage

func InitStorage() {
	if ErrorStorage != nil {
		return
	}

	ErrorStorage = &Storage{
		errors: make([]ExpressionError, 0),
		mutex:  sync.Mutex{},
	}
}

func (s *Storage) SaveError(expression, endpoint, errorType string) {
	e := ExpressionError{
		Expression: expression,
		Endpoint:   endpoint,
		Frequency:  1,
		Type:       errorType,
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i := 0; i < len(s.errors); i++ {
		if s.errors[i].Expression == e.Expression && s.errors[i].Type == e.Type {
			s.errors[i].Frequency++
			return
		}
	}

	s.errors = append(s.errors, e)
}

func (s *Storage) GetErrors() []ExpressionError {
	return s.errors
}

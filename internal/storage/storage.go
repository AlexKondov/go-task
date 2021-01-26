package storage

import (
	"sync"
)

type ExpressionError struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Type       string `json:"type"`
}

type Storage struct {
	errors []ExpressionError
	mutex  sync.RWMutex
}

func New() *Storage {
	return &Storage{
		errors: make([]ExpressionError, 0),
		mutex:  sync.RWMutex{},
	}
}

func (s *Storage) SaveError(expression, endpoint, errorType string) error {
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
			return nil
		}
	}

	s.errors = append(s.errors, e)

	return nil
}

func (s *Storage) GetErrors() []ExpressionError {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.errors
}

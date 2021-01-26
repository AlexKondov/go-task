package server

import (
	"bytes"
	"encoding/json"
	"github.com/AlexKondov/go-task/internal/storage"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func prepareTestServer(r http.Handler) *httptest.Server {
	s := httptest.NewServer(r)
	return s
}

func Test_evaluateHandler(t *testing.T) {
	s := storage.New()
	logger := logrus.New()
	handler := NewRequestHandler(s, logger)
	router := NewRouter(handler)
	server := prepareTestServer(router)

	payload := []byte(`{"expression": "What is 5 plus 4?"}`)
	req, _ := http.NewRequest("POST", server.URL + "/evaluate", bytes.NewBuffer(payload))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	expected := struct{
		Result int `json:"result"`
	}{
		Result: 9,
	}
	expectedJSON, err := json.Marshal(expected)

	if err != nil {
		t.Fatalf("Expected result marshal failed: %s", err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(recorder.Body)

	if string(expectedJSON) != string(bodyBytes) {
		t.Fatalf("Evaluate handler response does not match expected: %s", string(bodyBytes))
	}
}

func Test_validateHandler(t *testing.T) {
	s := storage.New()
	logger := logrus.New()
	handler := NewRequestHandler(s, logger)
	router := NewRouter(handler)
	server := prepareTestServer(router)

	payload := []byte(`{"expression": "What is 5 plus 4?"}`)
	req, _ := http.NewRequest("POST", server.URL + "/validate", bytes.NewBuffer(payload))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	expected := struct{
		Valid bool `json:"valid"`
	}{
		Valid: true,
	}
	expectedJSON, err := json.Marshal(expected)

	if err != nil {
		t.Fatalf("Expected result marshal failed: %s", err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(recorder.Body)

	if string(expectedJSON) != string(bodyBytes) {
		t.Fatalf("Evaluate handler response does not match expected: %s", string(bodyBytes))
	}
}

func Test_errorsHandler(t *testing.T) {
	s := storage.New()
	logger := logrus.New()
	handler := NewRequestHandler(s, logger)
	router := NewRouter(handler)
	server := prepareTestServer(router)

	payload := []byte(`{"expression": "What is 5 plus 4"}`)
	req, _ := http.NewRequest("POST", server.URL + "/evaluate", bytes.NewBuffer(payload))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	req, _ = http.NewRequest("GET", server.URL + "/errors", nil)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	expected := []storage.ExpressionError{
			{"What is 5 plus 4", "/evaluate", 1, "evaluation not terminated properly"},
	}
	expectedJSON, err := json.Marshal(expected)

	if err != nil {
		t.Fatalf("Expected result marshal failed: %s", err.Error())
	}

	bodyBytes, err := ioutil.ReadAll(recorder.Body)

	if string(expectedJSON) != string(bodyBytes) {
		t.Fatalf("Evaluate handler response does not match expected: %s", string(bodyBytes))
	}
}
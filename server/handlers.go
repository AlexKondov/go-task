package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/AlexKondov/go-task/api"
	"github.com/AlexKondov/go-task/internal/storage"
)

type RequestHandler struct {
	storage *storage.Storage
	logger  *logrus.Logger
}

func NewRequestHandler(storage *storage.Storage, logger *logrus.Logger) *RequestHandler {
	return &RequestHandler{
		storage: storage,
		logger:  logger,
	}
}

func (rh *RequestHandler) evaluate(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.logger.Error("bad request body")
		return
	}

	expression, ok := body["expression"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		rh.logger.Error("missing expression in request")
		return
	}

	num, err := api.EvaluateExpression(expression)

	response := make(map[string]interface{})
	response["result"] = num

	if err != nil {
		rh.logger.Errorf("Error occurred in expression: %s", expression)
		rh.storage.SaveError(expression, r.URL.Path, err.Error())
		response["error"] = err.Error()
	}

	resp, err := json.Marshal(response)

	if err != nil {
		rh.logger.Error("Error marshalling response")
	}

	w.Write(resp)
}

func (rh *RequestHandler) validate(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rh.logger.Fatal("Bad request body")
		return
	}

	valid, err := api.IsValidExpression(body["expression"])

	response := make(map[string]interface{})
	response["valid"] = valid

	if err != nil {
		rh.logger.Errorf("Error occurred in expression: %s", body["expression"])
		rh.storage.SaveError(body["expression"], r.URL.Path, err.Error())
		response["reason"] = err.Error()
	}

	resp, err := json.Marshal(response)

	if err != nil {
		rh.logger.Errorf("Error marshalling response")
	}

	w.Write(resp)
}

func (rh *RequestHandler) errors(w http.ResponseWriter, r *http.Request) {
	errors := api.GetExpressionErrors(rh.storage)

	resp, err := json.Marshal(errors)

	if err != nil {
		rh.logger.Errorf("Error marshalling response")
	}

	w.Write(resp)
}

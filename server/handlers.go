package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexKondov/go-task/api"
	"github.com/AlexKondov/go-task/internal/storage"
)

func evaluateHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Bad request body")
		return
	}

	num, err := api.EvaluateExpression(body["expression"])

	response := make(map[string]interface{})
	response["result"] = num

	if err != nil {
		log.Printf("Error occurred in expression: %s", body["expression"])
		storage.ErrorStorage.SaveError(body["expression"], r.URL.Path, err.Error())
		response["error"] = err.Error()
	}

	resp, err := json.Marshal(response)

	if err != nil {
		log.Printf("Error marshalling response")
	}

	w.Write(resp)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Bad request body")
		return
	}

	valid, err := api.IsValidExpression(body["expression"])

	response := make(map[string]interface{})
	response["valid"] = valid

	if err != nil {
		log.Printf("Error occurred in expression: %s", body["expression"])
		storage.ErrorStorage.SaveError(body["expression"], r.URL.Path, err.Error())
		response["reason"] = err.Error()
	}

	resp, err := json.Marshal(response)

	if err != nil {
		log.Printf("Error marshalling response")
	}

	w.Write(resp)
}

func errorsHandler(w http.ResponseWriter, r *http.Request) {
	errors := api.GetExpressionErrors()

	resp, err := json.Marshal(errors)

	if err != nil {
		log.Printf("Error marshalling response")
	}

	w.Write(resp)
}

package server

import (
	"encoding/json"
	"fmt"
	"go-task/api"
	"go-task/storage"
	"log"
	"net/http"
)

func evaluateHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		log.Fatal("Bad response")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	num, err := api.EvaluateExpression(body["expression"])

	response := make(map[string]interface{})
	response["result"] = num

	if err != nil {
		storage.ErrorStorage.SaveError(body["expression"], r.URL.Path, err.Error())
		response["error"] = err.Error()
	}

	resp, err := json.Marshal(response)

	if err != nil {
		fmt.Print("Test")
	}

	w.Write(resp)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	valid, err := api.IsValidExpression(body["expression"])

	response := make(map[string]interface{})
	response["valid"] = valid

	if err != nil {
		storage.ErrorStorage.SaveError(body["expression"], r.URL.Path, err.Error())
		response["reason"] = err.Error()
	}

	resp, err := json.Marshal(response)

	if err != nil {
		fmt.Print("Test")
	}

	w.Write(resp)
}

func errorsHandler(w http.ResponseWriter, r *http.Request) {
	errors := api.GetExpressionErrors()

	resp, err := json.Marshal(errors)

	if err != nil {
		fmt.Print("Test")
	}

	w.Write(resp)
}

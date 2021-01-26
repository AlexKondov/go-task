package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() error {
	r := mux.NewRouter()

	r.HandleFunc("/evaluate", evaluateHandler).Methods("POST")
	r.HandleFunc("/validate", validateHandler).Methods("POST")
	r.HandleFunc("/errors", errorsHandler).Methods("GET")

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", r)

	return err
}

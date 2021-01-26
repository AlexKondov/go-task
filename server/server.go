package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handler *RequestHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/evaluate", handler.evaluate).Methods("POST")
	r.HandleFunc("/validate", handler.validate).Methods("POST")
	r.HandleFunc("/errors", handler.errors).Methods("GET")

	return r
}

func StartServer(r http.Handler) error {
	err := http.ListenAndServe(":8080", r)
	return err
}

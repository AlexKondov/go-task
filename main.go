package main

import (
	"log"

	"github.com/AlexKondov/go-task/internal/storage"
	"github.com/AlexKondov/go-task/server"
)

func main() {
	storage.InitStorage()
	err := server.StartServer()

	if err == nil {
		log.Fatal("Error starting server")
	}
}

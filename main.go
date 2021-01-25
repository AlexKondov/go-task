package main

import (
	"go-task/server"
	"go-task/storage"
	"log"
)

func main() {
	storage.InitStorage()
	err := server.StartServer()

	if err == nil {
		log.Fatal("Error starting server")
	}
}

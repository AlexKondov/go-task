package main

import (
	"github.com/AlexKondov/go-task/internal/storage"
	"github.com/AlexKondov/go-task/server"
	"github.com/sirupsen/logrus"
)

func main() {
	storageInstance := storage.New()
	logger := logrus.New()

	handler := server.NewRequestHandler(storageInstance, logger)
	router := server.NewRouter(handler)

	logger.Info("Starting server on port 8080")
	err := server.StartServer(router)

	if err != nil {
		logger.Fatal("Error starting server")
	}
}

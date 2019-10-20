package main

import (
	"log"
	"os"

	"alice-echo/handler"
)

type Context struct {
	Handler *handler.Handler
}

func InitContext() (context *Context, err error) {
	secretPath := getSecretPath()

	echoHandler := handler.CreateEchoHandler()

	handlers := []handler.QuestionHandler {
		echoHandler,
	}

	handler, err := handler.CreateHandler(secretPath, handlers)

	if err != nil {
		log.Fatalf("Error while creating handler: %s", handler)
	}

	context = &Context{
		Handler: handler,
	}

	return
}

func getSecretPath() string {
	secretPath := os.Getenv("SECRET_PATH")
	if secretPath != "" {
		return secretPath
	} else {
		return "/webhook"
	}
}
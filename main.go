package main

import (
	"log"
	"net/http"
)

func main() {
	context, err := InitContext()

	if err != nil {
		log.Fatalf("Error while initializing context %s", err)
	}

	http.HandleFunc("/", context.Handler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
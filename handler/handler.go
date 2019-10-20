package handler

import (
	"log"
	"encoding/json"
	"net/http"

	"alice-echo/alice"
)

type Handler struct {
	secretPath string
	handlers []QuestionHandler
}

type QuestionHandler interface {
	Handle(question *alice.Question) (answer *alice.Answer, err error)
}

func CreateHandler(secretPath string, handlers []QuestionHandler) (handler *Handler, err error) {
	handler = &Handler {
		secretPath: secretPath,
		handlers: handlers,
	}

	return
}

func (this Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL == nil || r.URL.Path != this.secretPath {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var question alice.Question
	err := json.NewDecoder(r.Body).Decode(&question)

	if err != nil {
		log.Printf("Error while decoding request: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%#v", question)

	for _, handler := range this.handlers {
		resp, err := handler.Handle(&question)

		if err != nil {
			log.Printf("Error while processing handler: %s", err)
		}

		if resp != nil {
			log.Printf("Encoding response: %#v", resp)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			break
		}
	}
	
	

}
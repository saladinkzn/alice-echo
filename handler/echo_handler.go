package handler

import (
	"alice-echo/alice"
)

type EchoHandler struct {

}

func CreateEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (this EchoHandler) Handle(question *alice.Question) (answer *alice.Answer, err error) {
	response := alice.Response {
		Text: question.Request.OriginalUtterance,
		TTS: question.Request.OriginalUtterance,
		EndSession: false,
		Buttons: make([]alice.Button, 0),
	}

	answer = &alice.Answer {
		Version: question.Version,
		Session: convertSession(question.Session),
		Response: response,
	}

	return
}

func convertSession(session alice.Session) alice.Session {
	return alice.Session {
		SessionID: session.SessionID,
		MessageID: session.MessageID,
		UserID: session.UserID,
	}
}
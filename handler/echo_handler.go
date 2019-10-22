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
	originalUtterance := question.Request.OriginalUtterance
	var text string
	if len(originalUtterance) == 0 {
		text = "Пожалуйста, введите текст"
	} else {
		text = question.Request.OriginalUtterance
	}

	answer = makeAnswer(text, question.Version, question.Session)
	return
}

func makeAnswer(text string, version string, session alice.Session) *alice.Answer {
	response := alice.Response {
		Text: text,
		TTS: text,
		EndSession: false,
		Buttons: make([]alice.Button, 0),
	}

	answer := &alice.Answer {
		Version: version,
		Session: convertSession(session),
		Response: response,
	}

	return answer
}

func convertSession(session alice.Session) alice.Session {
	return alice.Session {
		SessionID: session.SessionID,
		MessageID: session.MessageID,
		UserID: session.UserID,
	}
}
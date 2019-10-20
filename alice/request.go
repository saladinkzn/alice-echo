package alice

type (
	Question struct {
		Meta Meta `json:"meta"`
		Request Request `json:"request"`
		Session Session `json:"session"`
		Version string `json:"version"`
	}

	Answer struct {
		Version string `json:"version"`
		Session Session `json:"session"`
		Response Response `json:"response"`
	}

	Response struct {
		Text string `json:"text"`
		TTS string `json:"tts"`
		Buttons []Button `json:"buttons"`
		EndSession bool `json:"end_session"`
	}

	Button struct {
		Title string `json:"title"`

		Payload Payload `json:"payload,omitempty"`

		URL string `json:"url"`

		Hide bool `json:"hide"`
	}

	Meta struct {
		Locale string `json:"locale"`

		TimeZone string `json:"timezone"`

		ClientID string `json:"client_id"`
	}

	Request struct {
		Type string `json:"type"`
		
		Markup Markup `json:"markup,omitempty"`

		Command string `json:"command"`

		OriginalUtterance string `json:"original_utterance"`

		Payload Payload `json:"payload,omitempty"`
	}

	Markup struct {
		DangerousContext bool `json:"dangerous_context,omitempy"`
	}

	Session struct {
		New bool `json:"new"`
		SessionID string `json:"session_id"`
		MessageID int64 `json:"message_id"`
		SkillID string `json:"skill_id"`
		UserID string `json:"user_id"`
	}

	Payload interface{}
)

const (
	SimpleUtterance = "SimpleUterrance"
	ButtonPressed = "ButtonPressed"
) 

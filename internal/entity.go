package internal

// HTTPRequest describes the request structure
type HTTPRequest struct {
	Meta struct {
		Locale     string `json:"locale"`
		Timezone   string `json:"timezone"`
		ClientId   string `json:"client_id"`
		Interfaces struct {
			Screen struct {
			} `json:"screen"`
			AccountLinking struct {
			} `json:"account_linking"`
			AudioPlayer struct {
			} `json:"audio_player"`
		} `json:"interfaces"`
	} `json:"meta"`
	Request struct {
		Type string `json:"type"`
	} `json:"request"`
	Session struct {
		MessageId int    `json:"message_id"`
		SessionId string `json:"session_id"`
		SkillId   string `json:"skill_id"`
		UserId    string `json:"user_id"`
		User      struct {
			UserId      string `json:"user_id"`
			AccessToken string `json:"access_token"`
		} `json:"user"`
		Application struct {
			ApplicationId string `json:"application_id"`
		} `json:"application"`
		New bool `json:"new"`
	} `json:"session"`
	State struct {
		Session struct {
			Value int `json:"value"`
		} `json:"session"`
		User struct {
			Value int `json:"value"`
		} `json:"user"`
		Application struct {
			Value int `json:"value"`
		} `json:"application"`
	} `json:"state"`
	Version string `json:"version"`
}

// HTTPResponse describes the response structure
type HTTPResponse struct {
	Response     Response `json:"response"`
	SessionState struct {
		Value int `json:"value"`
	} `json:"session_state"`
	UserStateUpdate struct {
		Value int `json:"value"`
	} `json:"user_state_update"`
	ApplicationState struct {
		Value int `json:"value"`
	} `json:"application_state"`
	Analytics struct {
		Events []struct {
			Name  string `json:"name"`
			Value struct {
				Field       string `json:"field"`
				SecondField struct {
					ThirdField string `json:"third field"`
				} `json:"second field"`
			} `json:"value,omitempty"`
		} `json:"events"`
	} `json:"analytics"`
	Version string `json:"version"`
}

type Response struct {
	Text    string `json:"text"`
	Tts     string `json:"tts"`
	Card    *Card  `json:"card,omitempty"`
	Buttons []struct {
		Title   string `json:"title"`
		Payload struct {
		} `json:"payload"`
		Url  string `json:"url"`
		Hide bool   `json:"hide"`
	} `json:"buttons"`
	EndSession bool `json:"end_session"`
	Directives struct {
	} `json:"directives"`
}

type Card struct {
	Type string `json:"type"`
}

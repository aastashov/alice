package internal

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

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

type Response struct {
	Text string `json:"text"`
	Tts  string `json:"tts"`
	Card struct {
		Type string `json:"type"`
	} `json:"card"`
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

func (h *Handler) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	log := h.logger.With("component", "callback_handler")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if r.Body != nil {
			_ = r.Body.Close()
		}
	}()

	// Parse the request body
	var req HTTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error("failed to decode request", "error", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Info("received request", "request", req)

	resp := HTTPResponse{
		Response: Response{
			Text:       "Это твой первый навык на Яндекс.Диалогах!",
			EndSession: false,
		},
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error("failed to encode response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

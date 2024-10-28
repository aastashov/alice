package internal

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

const (
	yandexDialogVersion = "1.0"
)

type Handler struct {
	logger  *slog.Logger
	release string
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
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
			Text:       "Этот навык еще в разработке",
			EndSession: false,
		},
		Version: yandexDialogVersion,
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Error("failed to encode response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

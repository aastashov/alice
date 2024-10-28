package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/aastashov/alice/internal"
)

var (
	Release = "unknown"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})).With("release", Release)
	log := logger.With("component", "app")

	apiHandler := internal.New(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/alice/dialogs", apiHandler.CallbackHandler)

	log.Info("starting server", "port", 8080)
	if err := http.ListenAndServe(":8080", mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to start server", "error", err)
	}
}

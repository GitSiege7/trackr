package main

import (
	"fmt"
	"net/http"

	"github.com/GitSiege7/trackr/internal/handlers"
	"github.com/GitSiege7/trackr/internal/helpers"
)

func main() {
	queries, err := helpers.OpenDB()
	helpers.HandlerError(err, "Failed to open db")

	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	cfg := &handlers.Config{
		Queries: queries,
	}

	mux.HandleFunc("GET /api/sessions", cfg.HandlerGetSessions)
	mux.HandleFunc("GET /api/sessions/{sessionID}", cfg.HandlerGetSession)
	mux.HandleFunc("GET /api/trackers", cfg.HandlerGetTrackers)
	mux.HandleFunc("GET /api/trackers/{trackerID}/sessions", cfg.HandlerGetSessionsByTracker)
	mux.HandleFunc("GET /api/sessions/ongoing", cfg.HandlerGetOngoing)
	mux.HandleFunc("POST /api/trackers/{trackerID}/sessions", cfg.HandlerCreateSession)
	mux.HandleFunc("DELETE /api/sessions/{sessionID}", cfg.HandlerDeleteSession)
	mux.HandleFunc("PATCH /api/sessions/{sessionID}/end", cfg.HandlerEndSession)
	mux.HandleFunc("PATCH /api/sessions/{sessionID}/note", cfg.HandlerUpdateNote)

	fmt.Println("Listening...")
	server.ListenAndServe()
}

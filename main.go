package main

import (
	"fmt"
	"net/http"

	"github.com/GitSiege7/trackr/database"
)

// Tests err, if not nil prints error output
func handlerError(err error, msg string) {
	if err != nil {
		fmt.Printf("ERROR: %v: %v\n", msg, err)
	}
}

type config struct {
	queries *database.Queries
}

func main() {
	queries, err := OpenDB()
	handlerError(err, "Failed to open db")

	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	cfg := &config{
		queries: queries,
	}

	mux.HandleFunc("GET /api/sessions", cfg.handlerGetSessions)
	mux.HandleFunc("GET /api/sessions/{sessionID}", cfg.handlerGetSession)
	mux.HandleFunc("GET /api/trackers", cfg.handlerGetTrackers)
	mux.HandleFunc("GET /api/trackers/{trackerID}/sessions", cfg.handlerGetSessionsByTracker)
	mux.HandleFunc("GET /api/sessions/ongoing", cfg.handlerGetOngoing)
	mux.HandleFunc("POST /api/trackers/{trackerID}/sessions", cfg.handlerCreateSession)
	mux.HandleFunc("DELETE /api/sessions/{sessionID}", cfg.handlerDeleteSession)
	mux.HandleFunc("PATCH /api/sessions/{sessionID}/end", cfg.handlerEndSession)
	mux.HandleFunc("PATCH /api/sessions/{sessionID}/note", cfg.handlerUpdateNote)

	fmt.Println("Listening...")
	server.ListenAndServe()
}

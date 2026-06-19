package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GitSiege7/trackr/database"
)

// Tests err, if not nil prints error output and exits
func handlerError(err error, msg string) {
	if err != nil {
		fmt.Printf("ERROR: %v: %v\n", msg, err)
		os.Exit(1)
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
	mux.HandleFunc("GET /api/trackers", cfg.handlerGetTrackers)
	mux.HandleFunc("GET /api/sessions/{trackerID}", cfg.handlerGetSessionsByID)
	mux.HandleFunc("GET /api/ongoing", cfg.handlerGetOngoing)
	// Start session (no end_datetime)
	// End session (add end_datetime)

	fmt.Println("Listening...")
	server.ListenAndServe()
}

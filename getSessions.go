package main

import (
	"context"
	"net/http"
)

func (cfg *config) handlerGetSessions(w http.ResponseWriter, r *http.Request) {
	sessions, err := cfg.queries.GetSessions(context.TODO())
	handlerError(err, "Failed to get sessions")

	err = respondWithJSON(w, 200, sessions)
	handlerError(err, "Failed to respond")
}

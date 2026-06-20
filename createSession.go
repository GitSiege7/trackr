package main

import (
	"context"
	"net/http"
	"strconv"
)

func (cfg *config) handlerCreateSession(w http.ResponseWriter, r *http.Request) {
	trackerID, err := strconv.ParseInt(r.PathValue("trackerID"), 10, 64)
	if err != nil {
		respondWithError(w, 400, "Failed to parse trackerID")
		return
	}

	session, err := cfg.queries.CreateSession(context.Background(), trackerID)
	if err != nil {
		respondWithError(w, 500, "Failed to create session")
		return
	}

	respondWithJSON(w, 201, session)
}

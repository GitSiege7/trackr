package main

import (
	"context"
	"net/http"
	"strconv"
)

func (cfg *config) handlerDeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		respondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	err = cfg.queries.DeleteSession(context.Background(), sessionID)
	if err != nil {
		respondWithError(w, 500, "Failed to delete session")
		return
	}

	w.WriteHeader(204)
}

package main

import (
	"context"
	"net/http"
	"strconv"
)

func (cfg *config) handlerGetSessionsByID(w http.ResponseWriter, r *http.Request) {
	trackerID, err := strconv.Atoi(r.PathValue("trackerID"))
	if err != nil {
		respondWithError(w, 400, "Failed to parse trackerID")
		return
	}

	sessions, err := cfg.queries.GetSessionsByID(context.TODO(), int64(trackerID))
	handlerError(err, "Failed to get sessions")

	err = respondWithJSON(w, 200, sessions)
	handlerError(err, "Failed to respond")
}

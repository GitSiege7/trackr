package internal

import (
	"context"
	"net/http"
	"strconv"
)

func (cfg *config) handlerGetSessionsByTracker(w http.ResponseWriter, r *http.Request) {
	trackerID, err := strconv.ParseInt(r.PathValue("trackerID"), 10, 64)
	if err != nil {
		respondWithError(w, 400, "Failed to parse trackerID")
		return
	}

	sessions, err := cfg.queries.GetSessionsByTracker(context.TODO(), trackerID)
	if err != nil {
		respondWithError(w, 500, "Failed to get sessions")
	}

	err = respondWithJSON(w, 200, sessions)
	handlerError(err, "Failed to respond")
}

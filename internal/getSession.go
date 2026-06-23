package internal

import (
	"context"
	"net/http"
	"strconv"
)

func (cfg *config) handlerGetSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		respondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	sessions, err := cfg.queries.GetSession(context.TODO(), sessionID)
	if err != nil {
		respondWithError(w, 500, "Failed to get sessions")
	}

	err = respondWithJSON(w, 200, sessions)
	handlerError(err, "Failed to respond")
}

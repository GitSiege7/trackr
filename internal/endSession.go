package internal

import (
	"context"
	"net/http"
	"strconv"
)

func (cfg *config) handlerEndSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		respondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	session, err := cfg.queries.EndSession(context.Background(), sessionID)
	if err != nil {
		respondWithError(w, 500, "Failed to update session")
		return
	}

	err = respondWithJSON(w, 200, session)
	handlerError(err, "Failed to respond")
}

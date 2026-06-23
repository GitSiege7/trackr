package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerEndSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	session, err := cfg.Queries.EndSession(context.Background(), sessionID)
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to update session")
		return
	}

	err = helpers.RespondWithJSON(w, 200, session)
	helpers.HandlerError(err, "Failed to respond")
}

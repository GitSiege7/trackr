package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerGetSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	sessions, err := cfg.Queries.GetSession(context.TODO(), sessionID)
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to get sessions")
	}

	err = helpers.RespondWithJSON(w, 200, sessions)
	helpers.HandlerError(err, "Failed to respond")
}

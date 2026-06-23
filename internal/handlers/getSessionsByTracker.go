package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerGetSessionsByTracker(w http.ResponseWriter, r *http.Request) {
	trackerID, err := strconv.ParseInt(r.PathValue("trackerID"), 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Failed to parse trackerID")
		return
	}

	sessions, err := cfg.Queries.GetSessionsByTracker(context.TODO(), trackerID)
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to get sessions")
	}

	err = helpers.RespondWithJSON(w, 200, sessions)
	helpers.HandlerError(err, "Failed to respond")
}

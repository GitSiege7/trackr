package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerCreateSession(w http.ResponseWriter, r *http.Request) {
	trackerID, err := strconv.ParseInt(r.PathValue("trackerID"), 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Failed to parse trackerID")
		return
	}

	session, err := cfg.Queries.CreateSession(context.Background(), trackerID)
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to create session")
		return
	}

	helpers.RespondWithJSON(w, 201, session)
}

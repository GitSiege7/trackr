package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerDeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	err = cfg.Queries.DeleteSession(context.Background(), sessionID)
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to delete session")
		return
	}

	w.WriteHeader(204)
}

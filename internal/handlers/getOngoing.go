package handlers

import (
	"context"
	"net/http"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerGetOngoing(w http.ResponseWriter, r *http.Request) {
	sessions, err := cfg.Queries.GetOngoingSessions(context.TODO())
	helpers.HandlerError(err, "Failed to get sessions")

	err = helpers.RespondWithJSON(w, 200, sessions)
	helpers.HandlerError(err, "Failed to respond")
}

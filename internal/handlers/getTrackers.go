package handlers

import (
	"context"
	"net/http"

	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerGetTrackers(w http.ResponseWriter, r *http.Request) {
	trackers, err := cfg.Queries.GetTrackers(context.TODO())
	helpers.HandlerError(err, "Failed to get trackers")

	err = helpers.RespondWithJSON(w, 200, trackers)
	helpers.HandlerError(err, "Failed to respond")
}

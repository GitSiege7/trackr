package main

import (
	"context"
	"net/http"
)

func (cfg *config) handlerGetTrackers(w http.ResponseWriter, r *http.Request) {
	trackers, err := cfg.queries.GetTrackers(context.TODO())
	handlerError(err, "Failed to get trackers")

	err = respondWithJSON(w, 200, trackers)
	handlerError(err, "Failed to respond")
}

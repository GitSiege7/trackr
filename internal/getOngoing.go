package internal

import (
	"context"
	"net/http"
)

func (cfg *config) handlerGetOngoing(w http.ResponseWriter, r *http.Request) {
	sessions, err := cfg.queries.GetOngoingSessions(context.TODO())
	handlerError(err, "Failed to get sessions")

	err = respondWithJSON(w, 200, sessions)
	handlerError(err, "Failed to respond")
}

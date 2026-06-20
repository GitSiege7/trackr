package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/database"
)

func (cfg *config) handlerUpdateNote(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		respondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	type req struct {
		Note string `json:"note"`
	}
	var dat req

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&dat)
	if err != nil {
		respondWithError(w, 500, "Failed to decode request body")
		return
	}

	session, err := cfg.queries.UpdateNote(context.Background(),
		database.UpdateNoteParams{
			Note: sql.NullString{
				String: dat.Note,
				Valid:  true,
			},
			ID: sessionID,
		})
	if err != nil {
		respondWithError(w, 500, "Failed to update session")
		return
	}

	err = respondWithJSON(w, 200, session)
	handlerError(err, "Failed to respond")
}

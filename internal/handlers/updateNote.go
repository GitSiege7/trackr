package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GitSiege7/trackr/database"
	"github.com/GitSiege7/trackr/internal/helpers"
)

func (cfg *Config) HandlerUpdateNote(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(r.PathValue("sessionID"), 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Failed to parse sessionID")
		return
	}

	type req struct {
		Note string `json:"note"`
	}
	var dat req

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&dat)
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to decode request body")
		return
	}

	session, err := cfg.Queries.UpdateNote(context.Background(),
		database.UpdateNoteParams{
			Note: sql.NullString{
				String: dat.Note,
				Valid:  true,
			},
			ID: sessionID,
		})
	if err != nil {
		helpers.RespondWithError(w, 500, "Failed to update session")
		return
	}

	err = helpers.RespondWithJSON(w, 200, session)
	helpers.HandlerError(err, "Failed to respond")
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DominikFeret/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (api *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error parsing json: ", err))
		return
	}

	users, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Error creating user: ", err))
		return
	}
}

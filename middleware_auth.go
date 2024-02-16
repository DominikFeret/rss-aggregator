package main

import (
	"fmt"
	"net/http"

	"github.com/DominikFeret/rss-aggregator/internal/auth"
	"github.com/DominikFeret/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprint("Error getting API key: ", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprint("Error getting user: ", err))
			return
		}

		handler(w, r, user)
	}
}

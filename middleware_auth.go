package main

import (
	"net/http"

	"github.com/rizkiromadoni/rss-aggregator/internal/auth"
	"github.com/rizkiromadoni/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, err.Error())
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, "Invalid API Key")
			return
		}

		handler(w, r, user)
	}
}

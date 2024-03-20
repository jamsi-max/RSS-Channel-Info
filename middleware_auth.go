package main

import (
	"fmt"
	"net/http"

	"github.com/jamsi-max/rss/internal/auth"
	"github.com/jamsi-max/rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/the-att-21/rssagg/internal/auth"
	"github.com/the-att-21/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCnf *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error : %v", err))
			return
		}

		user, err := apiCnf.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("User not found : %v", err))
			return
		}

		handler(w, r, user)
	}
}

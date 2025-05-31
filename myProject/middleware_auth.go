package main

import (
	"fmt"
	"net/http"

	"github.com/shubham-bhadra-10/golang_learning/myProject/internal/auth"
	"github.com/shubham-bhadra-10/golang_learning/myProject/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error getting API key: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error getting user: %v", err))
			return
		}
		handler(w, r, user)

	}
}

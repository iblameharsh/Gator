package main

import (
	"fmt"
	"net/http"

	"github.com/iblameharsh/Gator/internal/auth"
	"github.com/iblameharsh/Gator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error parsing json:", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Could not get the user:", err))
			return
		}

		handler(w, r, user)
	}
}

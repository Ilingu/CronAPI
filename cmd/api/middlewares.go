package main

import (
	"cron-api/cmd/utils"
	"net/http"
	"os"
)

type MiddlewareFunc func(http.Handler) http.Handler

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if utils.IsEmptyString(authToken) || utils.Hash(authToken) != os.Getenv("SERVER_KEY") {
			WriteResponse(&w, http.StatusUnauthorized, "Unauthorized Api Call")
			return
		}

		next.ServeHTTP(w, r)
	})
}

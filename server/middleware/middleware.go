package middleware

import (
	"net/http"

	"github.com/Backend-GoAPI-server/utils"
)

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	// Avoid middleware when you are going to login view
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/login" {
			next.ServeHTTP(rw, r)
		}
		if r.URL.Path == "/v1/signup" {
			next.ServeHTTP(rw, r)
		}

		jwtCookie, err := utils.GetTokenString(rw, r)
		utils.HandleErr(err)

		token, _, err := utils.ValidateToken(string(jwtCookie))
		utils.HandleErr(err)

		if !token.Valid {
			utils.ForbiddenException(rw)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

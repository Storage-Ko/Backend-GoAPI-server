package middleware

import (
	"net/http"

	"github.com/Backend-GoAPI-server/utils"
	"github.com/savsgio/go-logger/v2"
)

func AuthMiddleware(next http.Handler) http.Handler {
	// Avoid middleware when you are going to login view
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		if r.URL.Path == "/v1/login" {
			next.ServeHTTP(rw, r)
			return
		}
		if r.URL.Path == "/v1/signup" {
			next.ServeHTTP(rw, r)
			return
		}

		jwtCookie, err := utils.GetTokenString(rw, r)
		utils.HandleErr(err)

		token, user, err := utils.ValidateToken(string(jwtCookie))
		if err != nil {
			utils.UnauthorizedException(rw)
			return
		}

		logger.Info(user.Username)

		if !token.Valid {
			utils.UnauthorizedException(rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}

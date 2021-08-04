package middleware

import (
	"net/http"

	"github.com/Backend-GoAPI-server/utils"
	"github.com/savsgio/go-logger/v2"
)

// Auth Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Response content type setting
		rw.Header().Add("Content-Type", "application/json")

		// Exception Path of validation
		if r.URL.Path == "/v1/login" {
			next.ServeHTTP(rw, r)
			return
		}
		if r.URL.Path == "/v1/signup" {
			next.ServeHTTP(rw, r)
			return
		}

		// Get accessToken from request header
		jwtCookie, err := utils.GetTokenString(rw, r)
		utils.HandleErr(err)

		// Accesstoken validation
		token, user, err := utils.ValidateToken(string(jwtCookie))
		if err != nil {
			utils.UnauthorizedException(rw)
			return
		}

		// User data validation -- incomplete
		logger.Info(user.Username)

		// Accesstoken valid field validation
		if !token.Valid {
			utils.UnauthorizedException(rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}

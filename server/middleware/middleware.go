package middleware

import (
	"net/http"

	"github.com/Backend-GoAPI-server/utils"
	"github.com/savsgio/go-logger/v2"
)

// Exception Path of validation
func JSONResponseContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Response content type setting
		rw.Header().Add("Content-Type", "application/json")
		rw.Header().Add("Content-Type", "multipart/form-data")

		// Set CORS headers
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(rw, r)
	})
}

// Auth Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
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

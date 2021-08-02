package middleware

import (
	"net/http"
)

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

/*
func AuthMiddleware(ctx *atreugo.RequestCtx) error {
	// Avoid middleware when you are going to login view
	if string(ctx.Path()) == "/v1/login" {
		return ctx.Next()
	}
	if string(ctx.Path()) == "/v1/signup" {
		return ctx.Next()
	}

	jwtCookie, err := utils.GetTokenString(ctx)
	if err != nil {
		return err
	}

	token, _, err := utils.ValidateToken(string(jwtCookie))
	if err != nil {
		return err
	}

	if !token.Valid {
		return ctx.ErrorResponse(errors.New("your session is expired, login again please"), fasthttp.StatusForbidden)
	}

	return ctx.Next()
}
*/

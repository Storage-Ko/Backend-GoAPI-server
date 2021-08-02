package middleware

import (
	"errors"
	"net/http"

	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger/v2"
	"github.com/valyala/fasthttp"
)

func BeforeGlobal(ctx *atreugo.RequestCtx) error {
	logger.Info("Middleware executed BEFORE GLOBAL")

	return ctx.Next()
}

func AfterGlobal(ctx *atreugo.RequestCtx) error {
	logger.Info("Middleware executed AFTER GLOBAL")

	return ctx.Next()
}

func BeforeView(ctx *atreugo.RequestCtx) error {
	logger.Info("Middleware executed BEFORE VIEW")

	return ctx.Next()
}

func AfterView(ctx *atreugo.RequestCtx) error {
	logger.Info("Middleware executed AFTER VIEW")

	return ctx.Next()
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

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

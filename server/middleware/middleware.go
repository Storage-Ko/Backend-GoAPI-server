package middleware

import (
	"errors"
	"fmt"
	"strings"

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

func AuthMiddleware(ctx *atreugo.RequestCtx) error {
	// Avoid middleware when you are going to login view
	if string(ctx.Path()) == "/v1/login" {
		return ctx.Next()
	}

	buffer := ctx.Request.Header.String()
	slice := strings.Split(buffer, "Authorization: ")
	jwt := strings.Split(slice[1], "\nAccept: */*")

	jwtCookieStr := jwt[0]
	jwtCookie1 := []byte(jwtCookieStr)
	jwtCookie1 = jwtCookie1[:len(jwtCookie1)-1]
	fmt.Printf("AuthorizationA: %v\n", jwtCookie1)

	if len(jwtCookie1) == 0 {
		return ctx.ErrorResponse(errors.New("login required"), fasthttp.StatusForbidden)
	}

	token, _, err := utils.ValidateToken(string(jwtCookie1))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if !token.Valid {
		return ctx.ErrorResponse(errors.New("your session is expired, login again please"), fasthttp.StatusForbidden)
	}

	return ctx.Next()
}

package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/atreugo/v11"
	"github.com/savsgio/go-logger"
)

var jwtSignKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
	jwt.StandardClaims
}

func GetTokenString(ctx *atreugo.RequestCtx) ([]byte, error) {
	buffer := ctx.Request.Header.String()
	slice := strings.Split(buffer, "Authorization: ")
	jwt := strings.Split(slice[1], "\nAccept: */*")

	jwtCookieStr := jwt[0]
	jwtCookie := []byte(jwtCookieStr)
	jwtCookie = jwtCookie[:len(jwtCookie)-1]

	if len(jwtCookie) == 0 {
		ForbiddenException(ctx)
		return nil, errors.New("Forbidden Error")
	}
	return jwtCookie, nil
}

func GenerateToken(username []byte, password []byte) (string, time.Time) {
	logger.Debugf("Create new token for user %s", username)

	expireAt := time.Now().Add(1 * time.Minute)

	// Embed User information to `token`
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	})

	// token -> string. Only server knows the secret.
	tokenString, err := newToken.SignedString(jwtSignKey)
	if err != nil {
		logger.Error(err)
	}

	return tokenString, expireAt
}

func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	logger.Debug("Validating token...")

	user := &userCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	HandleErr(err)

	return token, user, err
}

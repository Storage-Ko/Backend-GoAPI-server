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
		return nil, errors.New("Token cannot found")
	}
	return jwtCookie, nil
}

func GenerateToken(username []byte, password []byte) string {
	logger.Debugf("Create new token for user %s", username)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	})

	tokenString, _ := newToken.SignedString(jwtSignKey)
	return tokenString
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

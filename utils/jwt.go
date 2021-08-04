package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/go-logger/v2"
)

var jwtSignKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetTokenString(rw http.ResponseWriter, r *http.Request) ([]byte, error) {
	jwt := r.Header.Get("Authorization")

	if len(jwt) == 0 {
		ForbiddenException(rw)
		return nil, errors.New("Token cannot found")
	}

	return []byte(jwt), nil
}

func AccessToken(username string) string {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	})

	access, err := accessToken.SignedString(jwtSignKey)
	HandleErr(err)

	return access
}

func RefreshToken(username string) string {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(),
		},
	})

	access, err := accessToken.SignedString(jwtSignKey)
	HandleErr(err)

	return access
}

func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	logger.Info("Validate Token")

	user := &userCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	HandleErr(err)

	return token, user, err
}

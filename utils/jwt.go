package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/go-logger"
)

var jwtSignKey = []byte("TestForFasthttpWithJWT")

type userCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
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

func GenerateToken(username []byte) string {
	logger.Debugf("Create new token for user %s", username)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	})

	tokenString, _ := newToken.SignedString(jwtSignKey)
	return tokenString
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

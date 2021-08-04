package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// test jwtSignKey
var jwtSignKey = []byte("TestForFasthttpWithJWT")

// Credential type
type userCredential struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GetToken function
func GetTokenString(rw http.ResponseWriter, r *http.Request) ([]byte, error) {
	// Get token from request token
	jwt := r.Header.Get("Authorization")

	// Token length validation
	if len(jwt) == 0 {
		UnauthorizedException(rw)
		return nil, errors.New("Token cannot found")
	}

	// Return token with type []byte
	return []byte(jwt), nil
}

// Generate accessToken
func AccessToken(username string) string {
	// Generate Token object
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(), // 10 Mins
		},
	})

	// Sign token
	access, err := accessToken.SignedString(jwtSignKey)
	HandleErr(err)

	return access
}

// Generate refreshToken
func RefreshToken(username string) string {
	// Generate Token object
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &userCredential{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(336 * time.Hour).Unix(), // 14Days
		},
	})

	// Sign token
	refresh, err := refreshToken.SignedString(jwtSignKey)
	HandleErr(err)

	return refresh
}

// Validate token
func ValidateToken(requestToken string) (*jwt.Token, *userCredential, error) {
	// Generate Credential object
	user := &userCredential{}

	// Parse token and validate
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})
	HandleErr(err)

	return token, user, err
}

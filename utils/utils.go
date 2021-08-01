package utils

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type userCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal("Error : " + err.Error())
		log.Panic(err)
	}
}

func GenerateToken(username string, password string) string {
	var jwtSecret map[string]string
	jwtSecret, err := godotenv.Read()
	HandleErr(err)
	var jwtSignKey = []byte(jwtSecret["JWT_SECRET"])
	expireAt := time.Now().Add(10 * time.Minute)

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
	HandleErr(err)
	return tokenString
}

func Hash(payload interface{}) string {
	s := fmt.Sprintf("%v", payload)
	hash := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func ByteToObj(payload []byte, object interface{}) {
	err := json.Unmarshal(payload, &object)
	HandleErr(err)
}

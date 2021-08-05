package utils

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
)

// Handle critical error with throw panic
func HandlePanic(err error) {
	if err != nil {
		logger.Error(err)
		log.Panic(err) // exit
	}
}

// Handle error
func HandleErr(err error) {
	if err != nil {
		logger.Error(err)
	}
}

// Hash payload & return hash string
func Hash(payload interface{}) string {
	s := fmt.Sprintf("%v", payload)
	hash := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// Translate payload with []byte type to object
func ByteToObj(payload []byte, object interface{}) {
	// payload type translate
	err := json.Unmarshal(payload, &object)
	if err != nil {
		logger.Error(err)
	}
}

func GetSecretKey() string {
	// jwt Key
	var jwtConfig map[string]string
	jwtConfig, err := godotenv.Read()
	HandleErr(err)

	var jwtSignKey = jwtConfig["JWT_SECRET"]
	return jwtSignKey
}

// Translate objest to byte & response payload with []byte type
func MarshalAndRW(status int, res interface{}, rw http.ResponseWriter) {
	// Set status code
	rw.WriteHeader(status)

	// Translate object to byte array
	resByte, err := json.MarshalIndent(res, "", "	")
	HandleErr(err)

	// Response payload
	rw.Write(resByte)
}

package utils

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/savsgio/go-logger/v2"
)

func HandleCredential(err error) {
	if err != nil {
		logger.Error(err)
		log.Panic(err)
	}
}

func HandleErr(err error) {
	if err != nil {
		logger.Error(err)
	}
}

func Hash(payload interface{}) string {
	s := fmt.Sprintf("%v", payload)
	hash := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func ByteToObj(payload []byte, object interface{}) {
	err := json.Unmarshal(payload, &object)
	if err != nil {
		logger.Error(err)
	}
}

func MarshalAndRW(status int, res interface{}, rw http.ResponseWriter) {
	rw.WriteHeader(status)
	logger.Info("1")
	resByte, err := json.MarshalIndent(res, "", "	")
	logger.Info("2")
	HandleErr(err)
	logger.Info("3")
	rw.Write(resByte)
	logger.Info("4")
}

package utils

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"

	"github.com/savsgio/go-logger/v2"
)

func HandleErr(err error) {
	if err != nil {
		logger.Error(err)
		log.Panic(err)
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

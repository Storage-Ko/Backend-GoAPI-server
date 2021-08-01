package utils

import (
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal("Error : " + err.Error())
		log.Panic(err)
	}
}

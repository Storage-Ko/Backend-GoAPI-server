package main

import (
	"github.com/Backend-GoAPI-server/db"
	"github.com/Backend-GoAPI-server/server"
)

func main() {
	db := db.Start()
	server.Start(4030)
	defer db.Close()
}

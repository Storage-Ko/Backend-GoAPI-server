package main

import (
	"github.com/Backend-GoAtreugo-server/db"
	"github.com/Backend-GoAtreugo-server/server"
)

func main() {
	db := db.Start()
	server.Start(4030)
	defer db.Close()
}

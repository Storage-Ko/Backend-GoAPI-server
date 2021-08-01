package server

import (
	"github.com/Backend-GoAtreugo-server/server/router"
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/savsgio/atreugo/v11"
)

func Start() {
	config := atreugo.Config{
		Addr: "0.0.0.0:4030",
	}
	server := atreugo.New(config)
	router.Start_v1(server)

	err := server.ListenAndServe()
	utils.HandleErr(err)
}

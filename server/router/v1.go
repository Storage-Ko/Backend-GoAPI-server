package router

import (
	"github.com/Backend-GoAtreugo-server/server/middleware"
	v1 "github.com/Backend-GoAtreugo-server/server/router/v1"
	"github.com/savsgio/atreugo/v11"
)

func Start_v1(server *atreugo.Atreugo) {
	version1 := server.NewGroupPath("/v1")
	version1.UseBefore(middleware.BeforeGlobal)
	version1.GET("/test", v1.TestHandle)
}

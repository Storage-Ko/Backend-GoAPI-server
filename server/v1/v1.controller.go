package v1

import (
	"github.com/Backend-GoAtreugo-server/server/middleware"
	"github.com/savsgio/atreugo/v11"
)

func Start_v1(server *atreugo.Atreugo) {
	version1 := server.NewGroupPath("/v1")
	version1.UseBefore(middleware.BeforeGlobal)
	version1.GET("/test", testHandle)
	version1.POST("/login", loginHandle)
}

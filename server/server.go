package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Backend-GoAPI-server/server/middleware"
	v1 "github.com/Backend-GoAPI-server/server/v1"
	"github.com/gorilla/mux"
	"github.com/savsgio/go-logger/v2"
)

var port string

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)

	router := mux.NewRouter()

	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.Use(middleware.AuthMiddleware)

	v1Router.HandleFunc("/document", v1.Documentation)
	v1Router.HandleFunc("/login", v1.LoginHandle)
	v1Router.HandleFunc("/signup", v1.SignupHandle)

	logger.Infof("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

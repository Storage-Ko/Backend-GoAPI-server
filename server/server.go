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

	router.Use(middleware.JsonContentTypeMiddleware)
	router.Use(middleware.AuthMiddleware)
	router.HandleFunc("/v1/document", v1.Documentation)
	router.HandleFunc("/v1/login", v1.LoginHandle)
	router.HandleFunc("/v1/signup", v1.SignupHandle)

	logger.Infof("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

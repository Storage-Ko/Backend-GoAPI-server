package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Backend-GoAPI-server/db"
	"github.com/Backend-GoAPI-server/server/file"
	"github.com/Backend-GoAPI-server/server/middleware"
	"github.com/Backend-GoAPI-server/server/user"
	v1 "github.com/Backend-GoAPI-server/server/v1"
	"github.com/gorilla/mux"
	"github.com/savsgio/go-logger/v2"
)

var port string

func Start(aPort int) {
	// Port setting
	port = fmt.Sprintf(":%d", aPort)

	// DB setting
	db.Start()
	defer db.CloseDB()
	logger.Info("Database is connected")

	db.Migrate()
	logger.Info("Migrating tables")

	// Main Router generate
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.JSONResponseContentType)

	router.HandleFunc("/file", file.UploadsHandler).Methods("POST")
	router.HandleFunc("/file/{path}", file.LoadsFile).Methods("GET")
	router.HandleFunc("/login", user.LoginHandle).Methods("POST")
	router.HandleFunc("/signup", user.SignupHandle).Methods("POST")

	// v1 SubRouter generate
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.Use(middleware.AuthMiddleware)

	// v1 Routes define
	v1Router.HandleFunc("/document", v1.Documentation).Methods("GET")
	v1Router.HandleFunc("/dropout/{id}", user.DropoutHandle).Methods("GET")
	v1Router.HandleFunc("/update", user.UpdateUserHandle).Methods("PUT")

	// Server Listen
	logger.Infof("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

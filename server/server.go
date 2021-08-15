package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Backend-GoAPI-server/db"
	"github.com/Backend-GoAPI-server/server/middleware"
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
	router.Use(middleware.JSONResponseContentType)

	router.Handle("/file", http.FileServer(http.Dir("public")))
	router.HandleFunc("/login", v1.LoginHandle).Methods("POST")
	router.HandleFunc("/signup", v1.SignupHandle).Methods("POST")
	router.HandleFunc("/upload", v1.UploadsHandler).Methods("POST")

	// v1 SubRouter generate
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.Use(middleware.AuthMiddleware)

	// v1 Routes define
	v1Router.HandleFunc("/document", v1.Documentation).Methods("GET")
	v1Router.HandleFunc("/dropout/{id}", v1.DropoutHandle).Methods("GET")
	v1Router.HandleFunc("/update", v1.UpdateUserHandle).Methods("PUT")

	// Server Listen
	logger.Infof("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

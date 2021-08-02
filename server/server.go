package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Backend-GoAtreugo-server/server/middleware"
	v1 "github.com/Backend-GoAtreugo-server/server/v1"
	"github.com/gorilla/mux"
)

var port string

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(middleware.JsonContentTypeMiddleware)
	router.HandleFunc("/v1", v1.Documentation).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

package server

import (
	"log"
	"net/http"

	handlers "github.com/WistleBlowers/totality-assignment-RESTful-version/handler"
)

func Start(host, port string) {

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5000"
	}
	addr := host + ":" + port
	server := http.Server{
		Addr:    addr,
		Handler: handlers.Routes(),
	}
	log.Println("server hosted at", server.Addr)
	log.Fatal(server.ListenAndServe())
}

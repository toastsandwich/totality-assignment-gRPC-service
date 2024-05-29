package server

import (
	"log"
	"net/http"

	handlers "github.com/toastsandwich/totality-assignment-GRPC-version/handler"
)

func Start(host, port string) error {
	addr := host + ":" + port
	server := http.Server{
		Addr:    addr,
		Handler: handlers.Routes(),
	}
	log.Println("server hosted at", server.Addr)
	return server.ListenAndServe()
}

package main

import (
	"os"

	"github.com/WistleBlowers/totality-assignment-RESTful-version/config"
	"github.com/WistleBlowers/totality-assignment-RESTful-version/server"
)

func main() {
	var host, port string
	if len(os.Args) == 1 {
		host = "localhost"
		port = "9000"
	} else {
		host, port = os.Args[1], os.Args[2]
	}
	config.Load()
	server.Start(host, port)
}

package main

import (
	"log"

	"github.com/WistleBlowers/totality-assignment-RESTful-version/config"
	"github.com/WistleBlowers/totality-assignment-RESTful-version/server"
)

func main() {
	config.Load()
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	server.Start(config.HOST, config.PORT)
}

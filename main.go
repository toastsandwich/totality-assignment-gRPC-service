package main

import (
	"log"

	"github.com/toastsandwich/totality-assignment-GRPC-version/config"
	"github.com/toastsandwich/totality-assignment-GRPC-version/server"
)

func main() {
	config.Load()

	file, err := config.InitLogger()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	err = config.ReadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = server.Start(config.HOST, config.PORT)
	if err != nil {
		log.Fatal(err.Error())
	}
}

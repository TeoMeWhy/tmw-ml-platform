package main

import (
	"log"
	"palantir/configs"
	"palantir/server"
)

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Deu merda!", err)
	}

	app, err := server.NewAppServer(config)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}

	app.App.Listen(":3000")

}

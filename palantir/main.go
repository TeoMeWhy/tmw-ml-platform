package main

import (
	"log"
	"palantir/configs"
	"palantir/server"
	"time"
)

func main() {

	time.Sleep(20 * time.Second) // Aguarda o MLflow Server iniciar

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

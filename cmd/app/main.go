package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/albakov/go-todo/internal/app/server"
)

func main() {
	config := server.NewConfig()

	_, err := toml.DecodeFile("configs/app.toml", config)
	if err != nil {
		log.Fatal()
	}

	server.Start(config)
}

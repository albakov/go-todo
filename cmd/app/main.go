package main

import (
	"github.com/albakov/go-todo/internal/app/server"
)

func main() {
	server.Start(&server.Config{
		BindAddr: ":8080",
	})
}

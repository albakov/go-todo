package server

import (
	"net/http"
	"time"
)

func Start(config *Config) {
	srv := apiServer()

	s := &http.Server{
		Handler:      srv,
		Addr:         config.BindAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.ListenAndServe()
}

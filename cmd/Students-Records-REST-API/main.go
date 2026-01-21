package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheSaifHub/Student-Records-REST-API/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my first go project"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Server failed to start.")
		}
	}()

}

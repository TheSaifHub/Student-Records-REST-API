package main

import (
	"context"
	// "fmt"

	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	slog.Info("Server Started", slog.String("Address", cfg.Addr))
	// fmt.Printf("Server started %s", cfg.Addr)

	// Graceful shutdown code begins here
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	err := server.ListenAndServe()
	// 	if err != nil && err != http.ErrServerClosed {
	// 		log.Fatal("Server failed to start.")
	// 	}
	// }()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", slog.String("error", err.Error()))
		}
	}()

	<-done

	slog.Info("Shutting down the server.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown Successfully.")

}

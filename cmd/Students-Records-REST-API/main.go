package main

import (
	"context"
	"log"

	// "fmt"

	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TheSaifHub/Student-Records-REST-API/internal/config"
	"github.com/TheSaifHub/Student-Records-REST-API/internal/http/handlers/student"
	"github.com/TheSaifHub/Student-Records-REST-API/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, error := sqlite.New(cfg)
	if error != nil {
		log.Fatal(error)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

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
			slog.Error("Server failed", slog.String("error", err.Error()))
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

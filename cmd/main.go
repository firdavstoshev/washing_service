package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/firdavstoshev/washing_service/api"
	"github.com/firdavstoshev/washing_service/api/handler"
	"github.com/firdavstoshev/washing_service/internal/service"
	"github.com/firdavstoshev/washing_service/internal/storage/postgres"
	"github.com/firdavstoshev/washing_service/pkg/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage := postgres.NewStorage(ctx, &cfg.Postgres)
	//storage.Migrate()
	defer storage.CloseDB()

	services := service.NewService(storage)
	handlers := handler.NewHandler(storage, services)
	router := api.SetupRoutes(handlers)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server is running on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}

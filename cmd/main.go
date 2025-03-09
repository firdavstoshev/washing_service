package main

import (
	"context"
	"net/http"
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
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage := postgres.NewStorage(ctx, &cfg.Postgres)
	services := service.NewService(storage)
	handlers := handler.NewHandler(storage, services)
	router := api.SetupRoutes(handlers)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	srv.ListenAndServe()
}

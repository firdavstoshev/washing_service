package postgres

import (
	"context"
	"fmt"

	"github.com/firdavstoshev/washing_service/internal/storage"
	"github.com/firdavstoshev/washing_service/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type store struct {
	order          *orderRepo
	washingService *washingServiceRepo
	customer       *customerRepo
}

func newStore(ctx context.Context, conn *gorm.DB) *store {
	return &store{
		order:          newOrderRepository(ctx, conn),
		washingService: newWashingServiceRepository(ctx, conn),
		customer:       newCustomerRepository(ctx, conn),
	}
}

func NewStorage(ctx context.Context, cfg *config.PostgresConfig) storage.IStorage {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Dushanbe",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return newStore(ctx, db)
}

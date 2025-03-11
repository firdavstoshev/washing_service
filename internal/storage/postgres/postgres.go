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
	conn *gorm.DB

	order          *orderRepo
	washingService *washingServiceRepo
	customer       *customerRepo
}

func newStore(ctx context.Context, conn *gorm.DB) *store {
	return &store{
		conn:           conn,
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

func (s *store) CloseDB() error {
	sqlDB, err := s.conn.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

const (
	unitTypesCreateQuery = `
CREATE TABLE IF NOT EXISTS unit_types -- Типы единиц измерения
(
    id   SERIAL PRIMARY KEY,  -- Уникальный идентификатор типа единицы измерения
    name VARCHAR(25) NOT NULL -- Название типа единицы измерения
);`

	serviceTypesCreateQuery = `
CREATE TABLE IF NOT EXISTS service_types
(
    id        SERIAL PRIMARY KEY,                                  -- Уникальный идентификатор типа услуги
    name      VARCHAR(255) NOT NULL,                               -- Название типа услуги
    unit_type INTEGER REFERENCES unit_types (id) ON DELETE CASCADE -- Тип единицы измерения (UnitType)
);`

	servicesCreateQuery = `
CREATE TABLE IF NOT EXISTS services
(
    id         SERIAL PRIMARY KEY,                                      -- Уникальный идентификатор услуги
    name       VARCHAR(255)   NOT NULL,                                 -- Название услуги
    type       INTEGER REFERENCES service_types (id) ON DELETE CASCADE, -- Тип услуги (ServiceType)
    unit_price DECIMAL(10, 2) NOT NULL                                  -- Цена за единицу (UnitPrice)
);`

	customersCreateQuery = `
CREATE TABLE IF NOT EXISTS customers
(
    id   SERIAL PRIMARY KEY,   -- Уникальный идентификатор клиента
    name VARCHAR(255) NOT NULL -- Имя клиента
);`

	ordersCreateQuery = `
CREATE TABLE IF NOT EXISTS orders
(
    id             SERIAL PRIMARY KEY,                                  -- Уникальный идентификатор заказа
    customer_id    INTEGER REFERENCES customers (id) ON DELETE CASCADE, -- Ссылка на клиента (FOREIGN KEY)
    is_child_items BOOLEAN        NOT NULL DEFAULT false,               -- Наличие детских вещей в заказе
    wait_days      INTEGER,                                             -- Количество дней, которые клиент готов подождать
    express        BOOLEAN                 DEFAULT FALSE,               -- Экспресс-услуга
    discount       DECIMAL(5, 2)           DEFAULT 0,                   -- Скидка на заказ в процентах
    total_price    DECIMAL(10, 2) NOT NULL                              -- Общая стоимость заказа
);`

	servicesItemsCreateQuery = `
CREATE TABLE IF NOT EXISTS service_items
(
    id         SERIAL PRIMARY KEY,                                 -- Уникальный идентификатор записи
    order_id   INTEGER REFERENCES orders (id) ON DELETE CASCADE,   -- Ссылка на заказ (FOREIGN KEY)
    service_id INTEGER REFERENCES services (id) ON DELETE CASCADE, -- Ссылка на услугу (FOREIGN KEY)
    amount     DECIMAL(10, 2) NOT NULL                             -- Количество (Amount)
);`
)

func (s *store) Migrate() error {
	createQueries := []string{unitTypesCreateQuery, serviceTypesCreateQuery, servicesCreateQuery, customersCreateQuery, ordersCreateQuery, servicesItemsCreateQuery}
	for i, query := range createQueries {
		if err := s.conn.Exec(query).Error; err != nil {
			return fmt.Errorf("failed to execute query %d: %w", i, err)
		}
	}
	return nil
}

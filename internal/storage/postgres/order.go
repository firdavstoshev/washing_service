package postgres

import (
	"context"

	"github.com/firdavstoshev/washing_service/internal/domain"
	"github.com/firdavstoshev/washing_service/internal/storage"

	"gorm.io/gorm"
)

type orderRepo struct {
	ctx  context.Context
	conn *gorm.DB
}

func newOrderRepository(ctx context.Context, conn *gorm.DB) *orderRepo {
	return &orderRepo{ctx: ctx, conn: conn}
}

func (s *store) Order() storage.IOrder {
	return s.order
}

// CreateOrderAndServiceItems - Создает заказ и связанные с ним записи о услугах
func (o *orderRepo) CreateOrderAndServiceItems(order *domain.Order, serviceItems *[]domain.ServiceItem) (uint, error) {
	tx := o.conn.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	err := tx.Create(order).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for i := range *serviceItems {
		(*serviceItems)[i].OrderID = order.ID
	}

	err = tx.Create(serviceItems).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit().Error
	if err != nil {
		return 0, err
	}

	return order.ID, nil
}

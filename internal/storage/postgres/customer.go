package postgres

import (
	"context"
	"errors"

	"github.com/firdavstoshev/washing_service/internal/domain"
	"github.com/firdavstoshev/washing_service/internal/storage"
	"github.com/firdavstoshev/washing_service/pkg/errs"

	"gorm.io/gorm"
)

type customerRepo struct {
	ctx  context.Context
	conn *gorm.DB
}

func newCustomerRepository(ctx context.Context, conn *gorm.DB) *customerRepo {
	return &customerRepo{ctx: ctx, conn: conn}
}

func (s *store) Customer() storage.ICustomer {
	return s.customer
}

func (c *customerRepo) GetCustomerByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := c.conn.First(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrCustomerNotFound
		}
		return nil, err
	}
	return &customer, nil
}

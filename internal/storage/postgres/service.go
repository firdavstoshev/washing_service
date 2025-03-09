package postgres

import (
	"context"
	"errors"

	"github.com/firdavstoshev/washing_service/internal/domain"
	"github.com/firdavstoshev/washing_service/internal/storage"
	"github.com/firdavstoshev/washing_service/pkg/errs"

	"gorm.io/gorm"
)

type serviceRepo struct {
	ctx  context.Context
	conn *gorm.DB
}

func newServiceRepository(ctx context.Context, conn *gorm.DB) *serviceRepo {
	return &serviceRepo{ctx: ctx, conn: conn}
}

func (s *store) Service() storage.IWashingService {
	return s.washingService
}

func (s *serviceRepo) GetWashingServices() ([]domain.Service, error) {
	var services []domain.Service
	if err := s.conn.Preload("Type").Preload("Type.UnitType").Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (s *serviceRepo) GetWashingServiceByID(id uint) (*domain.Service, error) {
	var service domain.Service
	if err := s.conn.Preload("Type").Preload("Type.UnitType").First(&service, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrWashingServiceNotFound
		}
		return nil, err
	}

	return &service, nil
}

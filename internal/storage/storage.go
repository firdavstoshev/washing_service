package storage

import "github.com/firdavstoshev/washing_service/internal/domain"

type IStorage interface {
	Migrate() error
	CloseDB() error

	Order() IOrder
	WashingService() IWashingService
	Customer() ICustomer
}

type IOrder interface {
	CreateOrderAndServiceItems(order *domain.Order, serviceItems *[]domain.ServiceItem) (uint, error)
}

type ICustomer interface {
	GetCustomerByID(id uint) (*domain.Customer, error)
}

type IWashingService interface {
	GetWashingServices() ([]domain.Service, error)
	GetWashingServiceByID(id uint) (*domain.Service, error)
}

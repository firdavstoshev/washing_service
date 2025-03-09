package service

import "github.com/firdavstoshev/washing_service/internal/storage"

type IService interface {
	Order() IOrderService
}

type service struct {
	orderService IOrderService
}

func NewService(storage storage.IStorage) IService {
	return &service{
		orderService: newOrderService(storage),
	}
}

func (s *service) Order() IOrderService {
	return s.orderService
}

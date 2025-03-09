package service

import (
	"log"

	"github.com/firdavstoshev/washing_service/internal/domain"
	"github.com/firdavstoshev/washing_service/internal/storage"
)

type IOrderService interface {
	CreateOrder(order *domain.Order, serviceItems *[]domain.ServiceItem) (uint, error)
	OrderPrice(order *domain.Order, serviceItems *[]domain.ServiceItem) (float64, error)
}

type orderService struct {
	storage storage.IStorage
}

func newOrderService(storage storage.IStorage) IOrderService {
	return &orderService{storage: storage}
}

func (o *orderService) OrderPrice(order *domain.Order, serviceItems *[]domain.ServiceItem) (float64, error) {
	customer, err := o.storage.Customer().GetCustomerByID(order.CustomerID)
	if err != nil {
		return 0, err
	}

	log.Printf("Customer: %v", customer)

	discount := 0.0

	// для детский вещей скидка 50%
	if order.IsChildItems {
		discount += 50
	}

	// если клиент ждет больше 5 дней скидка 30%
	if order.WaitDays > 5 {
		discount += 30
	}

	var (
		totalPrice  float64
		totalWeight float64
	)

	for _, item := range *serviceItems {
		svc, err := o.storage.Service().GetWashingServiceByID(item.ServiceID)
		if err != nil {
			return 0, err
		}

		totalPrice += svc.UnitPrice * item.Amount

		if svc.Type.UnitType.Name == "кг" { // TODO: use constant or enum
			totalWeight += svc.UnitPrice * item.Amount
		}
	}

	// если вес больше 10 кг скидка 20%
	if totalWeight > 10 {
		discount += 20
	}

	// если заказ экспресс 50% к цене
	if order.Express {
		totalPrice *= 1.5
	}

	// скидка не может быть больше 100%
	if discount > 100 {
		discount = 100
	}

	totalPrice -= totalPrice * discount / 100

	order.Discount = discount
	order.TotalPrice = totalPrice

	return totalPrice, nil
}

func (o *orderService) CreateOrder(order *domain.Order, serviceItems *[]domain.ServiceItem) (uint, error) {
	_, err := o.OrderPrice(order, serviceItems)
	if err != nil {
		return 0, err
	}

	orderId, err := o.storage.Order().CreateOrderAndServiceItems(order, serviceItems)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}

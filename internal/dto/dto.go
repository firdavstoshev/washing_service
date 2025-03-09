package dto

// UnitTypeDTO - Тип единицы измерения
type UnitTypeDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ServiceTypeDTO - Тип услуги
type ServiceTypeDTO struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	UnitType UnitTypeDTO `json:"unit_type"`
}

// ServiceDTO - Услуга
type ServiceDTO struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	TypeID    uint           `json:"type_id"`
	Type      ServiceTypeDTO `json:"type"`
	UnitPrice float64        `json:"unit_price"`
}

// CreateOrderRequest - Структура для создания заказа
type CreateOrderRequest struct {
	CustomerID   uint           `json:"customer_id" binding:"required"`
	IsChildItems bool           `json:"is_child_items"`
	WaitDays     int            `json:"wait_days"`
	Express      bool           `json:"express"`
	Services     []OrderService `json:"services" binding:"required"`
}

// CreateOrderResponse - Структура для ответа на запрос создания заказа
type CreateOrderResponse struct {
	OrderID uint `json:"order_id"`
}

// OrderService - Структура для представления услуги в заказе
type OrderService struct {
	ServiceID uint    `json:"service_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required,gt=0"`
}

type OrderPriceRequest struct {
	CustomerID   uint           `json:"customer_id" binding:"required"`
	Express      bool           `json:"express"`
	WaitDays     int            `json:"wait_days"`
	IsChildItems bool           `json:"is_child_items"`
	Services     []OrderService `json:"services" binding:"required"`
}

type OrderPriceResponse struct {
	Price float64 `json:"price"`
}

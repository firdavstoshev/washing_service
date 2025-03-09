package domain

// Order - Заказ
type Order struct {
	ID           uint     `gorm:"primaryKey;column:id"` // TODO: переделать на UUID
	CustomerID   uint     `gorm:"not null;column:customer_id"`
	Customer     Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE;column:customer_id"`
	IsChildItems bool     `gorm:"default:false;column:is_child_items"`
	WaitDays     int      `gorm:"default:0;column:wait_days"`
	Express      bool     `gorm:"default:false;column:express"`
	Discount     float64  `gorm:"type:decimal(5,2);default:0;column:discount"`
	TotalPrice   float64  `gorm:"type:decimal(10,2);default:0;column:total_price"`
}

func NewOrder(customerID uint, isChildItems, express bool, waitDays int) *Order {
	return &Order{
		CustomerID:   customerID,
		IsChildItems: isChildItems,
		WaitDays:     waitDays,
		Express:      express,
	}
}

func (o *Order) ApplyDiscount(discount float64) {
	o.Discount = discount
}

// ServiceItem - Запись о услуге в заказе
type ServiceItem struct {
	ID        uint    `gorm:"primaryKey;column:id"`
	OrderID   uint    `gorm:"not null;column:order_id"`
	ServiceID uint    `gorm:"not null;column:service_id"`
	Amount    float64 `gorm:"type:decimal(10,2);not null;column:amount"`
	Order     Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;column:order_id"`
	Service   Service `gorm:"foreignKey:ServiceID;constraint:OnDelete:CASCADE;column:service_id"`
}

package domain

// Customer - Клиент или посетитель
type Customer struct {
	ID   uint   `gorm:"primaryKey;column:id"`
	Name string `gorm:"type:varchar(255);not null;column:name"`
}

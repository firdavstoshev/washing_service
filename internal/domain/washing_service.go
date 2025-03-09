package domain

// UnitType - Тип единицы измерения
type UnitType struct {
	ID   uint   `gorm:"primaryKey;column:id"`
	Name string `gorm:"type:varchar(25);not null;column:name"`
}

// ServiceType - Тип услуги
type ServiceType struct {
	ID         uint     `gorm:"primaryKey;column:id"`
	Name       string   `gorm:"type:varchar(255);not null;column:name"`
	UnitTypeID uint     `gorm:"not null;column:unit_type"`
	UnitType   UnitType `gorm:"foreignKey:UnitTypeID;constraint:OnDelete:CASCADE;column:unit_type"`
}

// Service - Услуга
type Service struct {
	ID        uint        `gorm:"primaryKey;column:id"`
	Name      string      `gorm:"type:varchar(255);not null;column:name"`
	TypeID    uint        `gorm:"not null;column:type"`
	Type      ServiceType `gorm:"foreignKey:TypeID;constraint:OnDelete:CASCADE;column:type"`
	UnitPrice float64     `gorm:"type:decimal(10,2);not null;column:unit_price"`
}

package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID           uuid.UUID      `gorm:"column:id; primaryKey; not null; autoIncrement"`
	CustomerId   uuid.UUID      `gorm:"column:customer_id; not null;"`
	Address      string         `gorm:"column:address; not null"`
	PhoneNumber  string         `gorm:"column:phone_number; not null"`
	DeliveryDays time.Time      `gorm:"column:delivery_days; not null"`
	OrderDetails []OrderDetails `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment      string         `gorm:"column:payment; not null"`
	Total        float64        `gorm:"column:total; not null"`
	CreatedAt    time.Time      `gorm:"column:created_at; not null; autoCreateTime"`
}

func (Order) TableName() string {
	return "orders"
}

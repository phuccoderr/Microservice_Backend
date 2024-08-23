package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusCanceled  = "canceled"
)

type Order struct {
	ID           uuid.UUID      `gorm:"column:id; primaryKey; not null; autoIncrement"`
	CustomerId   string         `gorm:"column:customer_id; not null;"`
	Name         string         `gorm:"column:name; not null"`
	Address      string         `gorm:"column:address; not null"`
	PhoneNumber  string         `gorm:"column:phone_number; not null"`
	DeliveryDays time.Time      `gorm:"column:delivery_days; not null"`
	OrderDetails []OrderDetails `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment      string         `gorm:"column:payment; not null"`
	Total        float64        `gorm:"column:total; not null"`
	ProductCost  float64        `gorm:"column:product_cost; not null"`
	ShippingCost float64        `gorm:"column:shipping_cost; not null"`
	Status       string         `gorm:"column:status; not null"`
	CreatedAt    time.Time      `gorm:"column:created_at; not null; autoCreateTime"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}

func (Order) TableName() string {
	return "orders"
}

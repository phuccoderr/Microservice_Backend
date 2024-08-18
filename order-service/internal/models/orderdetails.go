package models

import (
	"github.com/google/uuid"
)

type OrderDetails struct {
	ID           uuid.UUID `gorm:"column:id; primaryKey; not null; autoIncrement"`
	ShippingCost float64   `gorm:"column:shipping_cost; not null;"`
	ProductCost  float64   `gorm:"column:product_cost; not null;"`
	ProductID    uuid.UUID `gorm:"column:product_id; not null"`
	Quantity     int64     `gorm:"column:quantity; not null"`
	OrderID      uuid.UUID `gorm:"column:order_id; not null"`
	TotalCost    float64   `gorm:"column:total_cost; not null"`
}

func (OrderDetails) TableName() string {
	return "order_details"
}

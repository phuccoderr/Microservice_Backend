package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetails struct {
	ID          uuid.UUID `gorm:"column:id; primaryKey; not null; autoIncrement"`
	ProductCost float64   `gorm:"column:product_cost; not null;"`
	ProductID   string    `gorm:"column:product_id; not null"`
	Quantity    int64     `gorm:"column:quantity; not null"`
	OrderID     uuid.UUID `gorm:"column:order_id; not null"`
	Total       float64   `gorm:"column:total; not null"`
}

func (o *OrderDetails) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}

func (OrderDetails) TableName() string {
	return "order_details"
}

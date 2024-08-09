package models

import (
	"time"
)

type Cart struct {
	Id         string    `gorm:"primary_key; column:id; type:uuid;"`
	CustomerId string    `gorm:"column:customer_id; not null"`
	ProductId  string    `gorm:"column:product_id; not null"`
	Quantity   int       `gorm:"column:quantity"`
	CreatedAt  time.Time `gorm:"column:created_at; autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at; autoUpdateTime"`
}

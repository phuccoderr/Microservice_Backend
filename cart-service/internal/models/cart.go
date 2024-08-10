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

type CartDto struct {
	Id         string    `json:"id"`
	CustomerId string    `json:"customer_id"`
	ProductId  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToCartDto(cart *Cart) CartDto {
	return CartDto{
		Id:         cart.Id,
		CustomerId: cart.CustomerId,
		ProductId:  cart.ProductId,
		Quantity:   cart.Quantity,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
	}
}

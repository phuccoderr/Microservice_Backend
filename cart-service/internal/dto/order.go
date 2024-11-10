package dto

import (
	"cart-service/internal/middleware"
	"time"
)

type PlaceOrderRequest struct {
	Address       string `json:"address" binding:"required"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	Note          string `json:"note"`
	Sale          int64  `json:"sale"`
}

type PlaceOrderMessage struct {
	CustomerId    CustomerDto `json:"customer_id"`
	Address       string      `json:"address"`
	PaymentMethod string      `json:"payment_method"`
	PhoneNumber   string      `json:"phone_number"`
	Total         float64     `json:"total"`
	ProductCost   float64     `json:"product_cost"`
	ShippingCost  float64     `json:"shipping_cost"`
	DeliverDays   time.Time   `json:"deliver_days"`
	OrderDetails  []CartDto   `json:"order_details"`
	Note          string      `json:"note"`
	CreatedAt     time.Time   `json:"created_at"`
}

func ToDtoPlaceOrderMessage(placeOrderMessage *PlaceOrderMessage, cartDto []CartDto, customerDto *middleware.CustomClaims) {

	placeOrderMessage.CustomerId.ID = customerDto.ID
	placeOrderMessage.CustomerId.Email = customerDto.Email
	placeOrderMessage.CustomerId.Name = customerDto.Name
	placeOrderMessage.OrderDetails = cartDto
}

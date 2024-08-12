package dto

import "time"

type CheckoutRequest struct {
	Address       string `json:"address" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type CheckoutDto struct {
	ProductTotal  float64   `json:"product_total"`
	ProductCost   float64   `json:"product_cost"`
	ShippingCost  float64   `json:"shipping_cost"`
	Address       string    `json:"address"`
	DeliverDays   time.Time `json:"deliver_days"`
	PaymentMethod string    `json:"payment_method"`
	CustomerEmail string    `json:"customer_email"`
}

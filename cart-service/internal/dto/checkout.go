package dto

import "time"

type CheckoutDto struct {
	ProductTotal float64   `json:"product_total"`
	ProductCost  float64   `json:"product_cost"`
	ShippingCost float64   `json:"shipping_cost"`
	DeliverDays  time.Time `json:"deliver_days"`
}

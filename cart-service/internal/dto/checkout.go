package dto

import "time"

type Checkout struct {
	ProductTotal float64   `json:"product_total"`
	ProductCost  float64   `json:"product_cost"`
	ShippingCost float64   `json:"shipping_cost"`
	Address      string    `json:"address"`
	DeliverDays  time.Time `json:"deliver_days"`
}

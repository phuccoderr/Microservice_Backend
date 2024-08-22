package kafka

import "time"

type VerifyCustomerMessage struct {
	Email     string `json:"email"`
	UrlVerify string `json:"url"`
}

type PlaceOrderMessage struct {
	CustomerId    string `json:"customer_id"`
	CustomerEmail string `json:"customer_email"`
	Address       string `json:"address"`
	PaymentMethod string `json:"payment_method"`
	PhoneNumber   string `json:"phone_number"`
	CheckOut      struct {
		ProductTotal float64   `json:"product_total"`
		ProductCost  float64   `json:"product_cost"`
		ShippingCost float64   `json:"shipping_cost"`
		DeliverDays  time.Time `json:"deliver_days"`
	} `json:"check_out"`
	Items []struct {
		ProductId    string  `json:"product_id"`
		ProductImage string  `json:"product_image"`
		Cost         float64 `json:"cost"`
		Price        float64 `json:"price"`
		Quantity     int64   `json:"quantity"`
		Total        float64 `json:"total"`
	} `json:"items"`
}

package kafka

import "time"

type PlaceOrderMessage struct {
	CustomerId struct {
		ID    string `json:"_id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"customer_id"`
	Address       string    `json:"address"`
	PaymentMethod string    `json:"payment_method"`
	PhoneNumber   string    `json:"phone_number"`
	Total         float64   `json:"total"`
	ProductCost   float64   `json:"product_cost"`
	ShippingCost  float64   `json:"shipping_cost"`
	DeliverDays   time.Time `json:"deliver_days"`
	OrderDetails  []struct {
		ProductId struct {
			Id          string  `json:"id"`
			Name        string  `json:"name"`
			Alias       string  `json:"alias"`
			Description string  `json:"description"`
			Status      bool    `json:"status"`
			Stock       int64   `json:"stock"`
			Cost        float64 `json:"cost"`
			Price       float64 `json:"price"`
			Sale        float64 `json:"sale"`
			ImageID     string  `json:"image_id"`
			URL         string  `json:"url"`

			ExtraImages []interface{} `json:"extra_images"`
			CategoryID  string        `json:"category_id"`
			CreatedAt   string        `json:"created_at"`
			UpdatedAt   string        `json:"updated_at"`
		} `json:"product_id"`
		Quantity int64 `json:"quantity"`
	} `json:"order_details"`
	Note string `json:"note"`
}

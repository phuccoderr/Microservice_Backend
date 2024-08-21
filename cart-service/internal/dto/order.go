package dto

type PlaceOrderRequest struct {
	Address       string `json:"address" binding:"required"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type PlaceOrderMessage struct {
	CustomerId    string       `json:"customer_id"`
	CustomerEmail string       `json:"customer_email"`
	Address       string       `json:"address"`
	PaymentMethod string       `json:"payment_method"`
	PhoneNumber   string       `json:"phone_number"`
	CheckOut      *CheckoutDto `json:"check_out"`
	Items         []CartDto    `json:"items"`
}

package dto

type PlaceOrder struct {
	Address       string `json:"address" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type PlaceOrderEvents struct {
	CustomerId    string        `json:"customer_id"`
	CustomerEmail string        `json:"customer_email"`
	Items         []CartRequest `json:"items"`
	Address       string        `json:"address"`
	PaymentMethod string        `json:"payment_method"`
	CheckOut      *CheckoutDto  `json:"check_out"`
}

package dto

type CartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}

type CartDto struct {
	CustomerId   string  `json:"customer_id"`
	ProductId    string  `json:"product_id"`
	ProductImage string  `json:"product_image"`
	Cost         float64 `json:"cost"`
	Price        float64 `json:"price"`
	Quantity     int64   `json:"quantity"`
	Total        float64 `json:"total"`
}

func ToCartDto(product *ProductResponse, customerId string) CartDto {
	return CartDto{
		CustomerId:   customerId,
		ProductId:    product.Id,
		ProductImage: product.URL,
		Cost:         product.Cost,
		Price:        product.Price,
		Quantity:     product.Quantity,
		Total:        product.Price * float64(product.Quantity),
	}
}

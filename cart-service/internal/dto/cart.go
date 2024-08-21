package dto

import "cart-service/pkg/response"

type CartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}

type CartDto struct {
	ProductId    string  `json:"product_id"`
	ProductImage string  `json:"product_image"`
	Cost         float64 `json:"cost"`
	Price        float64 `json:"price"`
	Quantity     int64   `json:"quantity"`
	Total        float64 `json:"total"`
}

func ToCartDto(product *response.ProductResponse, customerId string, quantity int64) CartDto {
	return CartDto{
		ProductId:    product.Id,
		ProductImage: product.URL,
		Cost:         product.Cost * float64(quantity),
		Price:        product.Price,
		Quantity:     quantity,
		Total:        product.Price * float64(quantity),
	}
}

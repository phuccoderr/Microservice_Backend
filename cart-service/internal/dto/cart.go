package dto

import (
	"cart-service/internal/middleware"
	"cart-service/pkg/response"
)

type CartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}

type CartDto struct {
	ProductId  *response.ProductResponse `json:"product_id"`
	CustomerId CustomerDto               `json:"customer_id"`
	Quantity   int64                     `json:"quantity"`
}

func ToCartDto(product *response.ProductResponse, customerDto *middleware.CustomClaims, quantity int64) CartDto {
	var cartDto CartDto
	cartDto.ProductId = product
	cartDto.CustomerId.ID = customerDto.ID
	cartDto.CustomerId.Name = customerDto.Name
	cartDto.CustomerId.Email = customerDto.Email
	cartDto.Quantity = quantity

	return cartDto
}

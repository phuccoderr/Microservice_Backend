package dto

import (
	"cart-service/internal/models"
	"time"
)

type CartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartDto struct {
	Id           string    `json:"id"`
	CustomerId   string    `json:"customer_id"`
	ProductId    string    `json:"product_id"`
	ProductImage string    `json:"product_image"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	Total        float64   `json:"total"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ToListCartDto(carts []models.Cart) []CartDto {
	var listCartDto []CartDto
	for _, item := range carts {
		dto := ToCartDto(&item)
		listCartDto = append(listCartDto, dto)
	}
	return listCartDto
}

func ToCartDto(cart *models.Cart) CartDto {
	return CartDto{
		Id:           cart.Id,
		CustomerId:   cart.CustomerId,
		ProductId:    cart.ProductId,
		ProductImage: cart.ProductImage,
		Price:        cart.Price,
		Quantity:     cart.Quantity,
		Total:        cart.Total,
		CreatedAt:    cart.CreatedAt,
		UpdatedAt:    cart.UpdatedAt,
	}
}

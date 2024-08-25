package dto

import (
	"github.com/google/uuid"
	"order-service/internal/models"
	"time"
)

type OrderDto struct {
	ID           uuid.UUID         `json:"id"`
	CustomerId   string            `json:"customer_id"`
	Name         string            `json:"name"`
	Address      string            `json:"address"`
	PhoneNumber  string            `json:"phone_number"`
	DeliveryDays time.Time         `json:"delivery_days"`
	OrderDetails []OrderDetailsDto `json:"order_details"`
	Payment      string            `json:"payment"`
	Total        float64           `json:"total"`
	ProductCost  float64           `json:"product_cost"`
	ShippingCost float64           `json:"shipping_cost"`
	Status       string            `json:"status"`
	CreatedAt    time.Time         `json:"created_at"`
}

type OrderDetailsDto struct {
	ID          uuid.UUID `json:"id"`
	ProductCost float64   `json:"product_cost"`
	ProductID   string    `json:"product_id"`
	Quantity    int64     `json:"quantity"`
	OrderID     uuid.UUID `json:"order_id"`
	Total       float64   `json:"total"`
}

func ListEntityToDto(orders []models.Order) []OrderDto {
	listDto := make([]OrderDto, 0, len(orders))
	for _, order := range orders {
		listDto = append(listDto, *EntityToDto(&order))
	}

	return listDto
}

func EntityToDto(order *models.Order) *OrderDto {
	orderDto := &OrderDto{
		ID:           order.ID,
		CustomerId:   order.CustomerId,
		Name:         order.Name,
		Address:      order.Address,
		PhoneNumber:  order.PhoneNumber,
		DeliveryDays: order.DeliveryDays,
		Payment:      order.Payment,
		Total:        order.Total,
		ProductCost:  order.ProductCost,
		ShippingCost: order.ShippingCost,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
	}

	for _, item := range order.OrderDetails {
		detail := OrderDetailsDto{
			ID:          item.ID,
			ProductCost: item.ProductCost,
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			OrderID:     item.OrderID,
			Total:       item.Total,
		}
		orderDto.OrderDetails = append(orderDto.OrderDetails, detail)
	}

	return orderDto
}

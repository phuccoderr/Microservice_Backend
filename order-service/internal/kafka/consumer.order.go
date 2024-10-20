package kafka

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"order-service/global"
	"order-service/internal/models"
	"order-service/internal/service"
)

func ConsumeOrder(service service.IOrderRedisService) {

	message, err := global.Consumer.ReadMessage(context.Background())
	if err != nil {
		global.Logger.Error("Error reading message", zap.Error(err))
		log.Fatalln(err)
	}
	var placeOrderMessage PlaceOrderMessage
	err = json.Unmarshal(message.Value, &placeOrderMessage)
	if err != nil {
		global.Logger.Error("Error unmarshalling message", zap.Error(err))
		log.Fatalln(err)
	}

	newOrder := models.Order{
		Name:         placeOrderMessage.CustomerId.Name,
		Address:      placeOrderMessage.Address,
		PhoneNumber:  placeOrderMessage.PhoneNumber,
		CustomerId:   placeOrderMessage.CustomerId.ID,
		Payment:      placeOrderMessage.PaymentMethod,
		ProductCost:  placeOrderMessage.ProductCost,
		ShippingCost: placeOrderMessage.ShippingCost,
		Total:        placeOrderMessage.Total,
		DeliveryDays: placeOrderMessage.DeliverDays,
		Status:       models.StatusPending,
		Note:         placeOrderMessage.Note,
	}

	for _, item := range placeOrderMessage.OrderDetails {
		var sale float64
		sale = item.ProductId.Price * (item.ProductId.Sale / 100)
		detail := models.OrderDetails{
			ProductCost: item.ProductId.Cost,
			Total:       (item.ProductId.Price - sale) * float64(item.Quantity),
			Quantity:    item.Quantity,
			ProductID:   item.ProductId.Id,
		}
		newOrder.OrderDetails = append(newOrder.OrderDetails, detail)
	}
	service.Clear("all_orders:*")
	global.Mdb.Create(&newOrder)
}

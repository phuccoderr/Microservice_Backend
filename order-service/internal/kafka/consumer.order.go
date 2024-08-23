package kafka

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"order-service/global"
	"order-service/internal/models"
)

func ConsumeOrder() {

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
		Name:         placeOrderMessage.CustomerName,
		Address:      placeOrderMessage.Address,
		PhoneNumber:  placeOrderMessage.PhoneNumber,
		CustomerId:   placeOrderMessage.CustomerId,
		Payment:      placeOrderMessage.PaymentMethod,
		ProductCost:  placeOrderMessage.CheckOut.ProductCost,
		ShippingCost: placeOrderMessage.CheckOut.ShippingCost,
		Total:        placeOrderMessage.CheckOut.ProductTotal,
		DeliveryDays: placeOrderMessage.CheckOut.DeliverDays,
		Status:       models.StatusPending,
	}

	for _, item := range placeOrderMessage.Items {
		detail := models.OrderDetails{
			ProductCost: item.Cost,
			Total:       item.Total,
			Quantity:    item.Quantity,
			ProductID:   item.ProductId,
		}
		newOrder.OrderDetails = append(newOrder.OrderDetails, detail)
	}
	global.Mdb.Create(&newOrder)
}

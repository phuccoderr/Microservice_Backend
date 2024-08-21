package kafka

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"log"
	"order-service/global"
	"order-service/internal/models"
	"time"
)

type PlaceOrderMessage struct {
	CustomerId    string `json:"customer_id"`
	CustomerEmail string `json:"customer_email"`
	Items         []struct {
		ProductId string `json:"product_id"`
		Quantity  int64  `json:"quantity"`
	} `json:"items"`
	Address       string `json:"address"`
	PaymentMethod string `json:"payment_method"`
	CheckOut      struct {
		ProductTotal float64   `json:"product_total"`
		ProductCost  float64   `json:"product_cost"`
		ShippingCost float64   `json:"shipping_cost"`
		DeliverDays  time.Time `json:"deliver_days"`
	} `json:"check_out"`
}

func ConsumeOrder() {
	defer global.Consumer.Close()

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

	order := models.Order{}
	order.Address = placeOrderMessage.Address
	order.PhoneNumber = "0123"
	order.DeliveryDays = placeOrderMessage.CheckOut.DeliverDays
	order.Payment = placeOrderMessage.PaymentMethod
	order.Total = placeOrderMessage.CheckOut.ProductTotal
	global.Mdb.Create(&order)
}

package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mail-service/global"
	"mail-service/internal/constants"
	"mail-service/internal/controller"
	"mail-service/internal/kafka"
	"mail-service/internal/middleware"
	"mail-service/internal/service"
)

func NewRouter(r *gin.Engine) {
	go func() {
		defer global.Consume.Close()
		for {
			message, err := global.Consume.ReadMessage(context.Background())
			if err != nil {
				global.Logger.Error("Error reading message", zap.Error(err))
				continue
			}

			switch message.Topic {
			case constants.KAFKA_TOPIC_VERIFY_ACCOUNT:
				kafka.HandleVerifyCustomer(&message)
			case constants.KAFKA_TOPIC_VERIFY_PASSWORD:
				kafka.HandleVerifyPassword(&message)
			case constants.KAFKA_TOPIC_PLACE_ORDER:
				kafka.HandlePlaceOrderCustomer(&message)
			default:
				global.Logger.Info("unknown topic", zap.String("topic", message.Topic))
			}
		}
	}()

	r.Use(middleware.CorsMiddleware())

	mailService := service.NewMailService()
	mailController := controller.NewMailController(mailService)

	r.POST("/api/v1/feedback", mailController.SendFeedBack)

}

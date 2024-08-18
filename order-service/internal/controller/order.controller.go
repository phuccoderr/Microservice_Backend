package controller

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/service"
	"order-service/pkg/response"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderService: service.NewOrderService(),
	}
}

func (oc *OrderController) GetOrderList(c *gin.Context) {
	response.SuccessResponse(c, 200, oc.orderService.GetALL())
}

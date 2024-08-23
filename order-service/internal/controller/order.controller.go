package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/constants"
	"order-service/internal/service"
	"order-service/pkg/response"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "asc")
	keyword := c.DefaultQuery("keyword", "")

	orders, err := oc.orderService.ListByPage(page, limit, sort, keyword)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_STATUS_UNPROCESSABLEENTITTY, http.StatusUnprocessableEntity)
		return
	}

	response.SuccessResponse(c, http.StatusOK, orders)
}

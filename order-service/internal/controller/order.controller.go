package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/constants"
	"order-service/internal/dto"
	"order-service/internal/middlewares"
	"order-service/internal/service"
	"order-service/pkg/response"
	"strconv"
)

type OrderController struct {
	orderService      service.IOrderService
	orderRedisService service.IOrderRedisService
}

func NewOrderController(orderService service.IOrderService, redisService service.IOrderRedisService) *OrderController {
	return &OrderController{
		orderService:      orderService,
		orderRedisService: redisService,
	}
}

func (oc *OrderController) GetAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "asc")
	keyword := c.DefaultQuery("keyword", "")

	ordersRedis, err := oc.orderRedisService.GetAllOrder(page, limit, sort)
	if err == nil {
		response.SuccessResponse(c, http.StatusOK, ordersRedis)
		return
	}

	orders, err := oc.orderService.ListByPage(page, limit, sort, keyword)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_STATUS_UNPROCESSABLEENTITTY, http.StatusUnprocessableEntity)
		return
	}

	paginationDto := dto.BuildPaginationDto(orders, page, limit)
	oc.orderRedisService.SetOrder(paginationDto, page, limit, sort)

	response.SuccessResponse(c, http.StatusOK, paginationDto)
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := oc.orderService.FindOne(id)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound)
		return
	}

	response.SuccessResponse(c, http.StatusOK, dto.EntityToDto(order))

}

func (oc *OrderController) ChangeStatus(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	err := oc.orderService.UpdateStatus(id, status)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound)
		return
	}

	response.SuccessResponse(c, http.StatusOK, "")
}

func (oc *OrderController) GetOrderByCustomer(c *gin.Context) {
	_, customer, err := middlewares.JWTGetTokenAndCustomer(c)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized)
		return
	}

	orders, err := oc.orderService.FindByCustomer(customer.ID)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound)
		return
	}

	ordersDto := dto.ListEntityToDto(orders)
	response.SuccessResponse(c, http.StatusOK, ordersDto)
}

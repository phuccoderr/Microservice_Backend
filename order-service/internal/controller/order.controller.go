package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"order-service/global"
	"order-service/internal/cache"
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
	productService    service.IProductService
}

func NewOrderController(orderService service.IOrderService, redisService service.IOrderRedisService, productService service.IProductService) *OrderController {
	return &OrderController{
		orderService:      orderService,
		orderRedisService: redisService,
		productService:    productService,
	}
}

func (oc *OrderController) GetAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "asc")
	keyword := c.DefaultQuery("keyword", "")

	ordersRedis, err := oc.orderRedisService.GetAllOrder(cache.OrdersKey(page, limit, sort))
	if err == nil {
		response.SuccessResponse(c, http.StatusOK, ordersRedis)
		return
	}

	orders, err := oc.orderService.ListByPage(page, limit, sort, keyword)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_STATUS_UNPROCESSABLEENTITTY, http.StatusUnprocessableEntity)
		return
	}

	ordersDto := dto.ListEntityToDto(orders)
	for i, orderDto := range ordersDto {
		for j, detail := range orderDto.OrderDetails {
			product, err := oc.productService.GetProductById(detail.ProductID)
			if err != nil {
				global.Logger.Error("Get Product Error", zap.Error(err))
				response.ErrorResponse(c, err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError)
				return
			}
			ordersDto[i].OrderDetails[j].Product = product
		}
	}

	paginationDto := dto.BuildPaginationDto(ordersDto, page, limit)
	oc.orderRedisService.SetOrder(paginationDto, cache.OrdersKey(page, limit, sort))

	response.SuccessResponse(c, http.StatusOK, paginationDto)
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := oc.orderService.FindOne(id)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound)
		return
	}

	orderDto := dto.EntityToDto(order)

	for i, detail := range orderDto.OrderDetails {
		product, err := oc.productService.GetProductById(detail.ProductID)
		if err != nil {
			global.Logger.Error("Get Product Error", zap.Error(err))
			response.ErrorResponse(c, err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError)
			return
		}
		orderDto.OrderDetails[i].Product = product
	}

	response.SuccessResponse(c, http.StatusOK, orderDto)

}

func (oc *OrderController) ChangeStatus(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	err := oc.orderService.UpdateStatus(id, status)
	if err != nil {
		response.ErrorResponse(c, err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound)
		return
	}

	oc.orderRedisService.Clear("all_orders:*")
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
	for i, orderDto := range ordersDto {
		for j, detail := range orderDto.OrderDetails {
			product, err := oc.productService.GetProductById(detail.ProductID)
			if err != nil {
				global.Logger.Error("Get Product Error", zap.Error(err))
				response.ErrorResponse(c, err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError)
				return
			}
			ordersDto[i].OrderDetails[j].Product = product
		}
	}
	oc.orderRedisService.SetOrder(ordersDto, cache.OrdersByCustomerKey(customer.ID))

	response.SuccessResponse(c, http.StatusOK, ordersDto)
}

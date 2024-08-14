package cart

import (
	"cart-service/internal/constants"
	"cart-service/internal/controller/microservices"
	"cart-service/internal/dto"
	"cart-service/internal/kafka"
	"cart-service/internal/middleware"
	"cart-service/internal/response"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cartService ICartService
}

func NewHandler(cartService ICartService) *Handler {
	return &Handler{cartService: cartService}
}

func (h *Handler) AddProductToCart(c *gin.Context) {
	var cartRequest *dto.CartRequest
	err := c.ShouldBindJSON(&cartRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.BuildResponseError(err.Error(), constants.STATUS_BAD_REQUEST, http.StatusBadRequest))
		return
	}

	customer, err := middleware.JWTGetCustomer(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	newQuantity, err := h.cartService.addProductToCart(customer.ID, cartRequest.ProductId, cartRequest.Quantity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.BuildResponseError(err.Error(), constants.STATUS_BAD_REQUEST, http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusCreated, response.BuildResponseObject(constants.CREATE_SUCCESS, http.StatusCreated, newQuantity))

}

func (h *Handler) GetCart(c *gin.Context) {

	token, customer, err := middleware.JWTGetTokenAndCustomer(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
		return
	}

	cartsRedis, err := h.cartService.getCart(customer.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	var carts []dto.CartDto
	for _, cart := range cartsRedis {
		product, err := microservices.CallGetProduct(cart.ProductId, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
			return
		}
		carts = append(carts, dto.ToCartDto(product, customer.ID, cart.Quantity))
	}

	c.JSON(http.StatusOK, response.BuildResponseObject(constants.GET_SUCCESS, http.StatusOK, carts))

}

func (h *Handler) DeleteCart(c *gin.Context) {
	productId := c.Param("id")

	if productId == "" {
		c.JSON(http.StatusBadRequest,
			response.BuildResponseError("productId is required", constants.STATUS_BAD_REQUEST, http.StatusBadRequest))
		return
	}

	customer, err := middleware.JWTGetCustomer(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
		return
	}

	cart, err := h.cartService.checkProductInCart(customer.ID, productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	if !cart {
		c.AbortWithStatusJSON(http.StatusNotFound,
			response.BuildResponseError(constants.DB_NOT_FOUND, constants.STATUS_NOT_FOUND, http.StatusNotFound))
		return
	}

	err = h.cartService.deleteCart(customer.ID, productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.BuildResponseObject(constants.DELETE_SUCCESS, http.StatusOK, nil))
}

func (h *Handler) Checkout(c *gin.Context) {
	token, customer, err := middleware.JWTGetTokenAndCustomer(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
		return
	}

	cartsRedis, err := h.cartService.getCart(customer.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	var carts []dto.CartDto
	for _, cart := range cartsRedis {
		product, err := microservices.CallGetProduct(cart.ProductId, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
			return
		}
		if cart.Quantity > product.Stock {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				response.BuildResponseError(fmt.Sprintf("quantity less than %v", product.Stock), constants.STATUS_STATUS_UNPROCESSABLEENTITTY, http.StatusUnprocessableEntity))
			return
		}
		carts = append(carts, dto.ToCartDto(product, customer.ID, cart.Quantity))
	}
	if len(carts) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			response.BuildResponseError(constants.DB_NOT_FOUND, constants.STATUS_NOT_FOUND, http.StatusNotFound))
		return
	}

	checkOutInfo := h.cartService.checkout(carts, customer.Email)
	c.JSON(http.StatusOK, response.BuildResponseObject(constants.CHECKOUT_SUCCESS, http.StatusOK, checkOutInfo))

}

func (h *Handler) PlaceOrder(c *gin.Context) {
	placeOrder := &dto.PlaceOrder{}
	err := c.ShouldBindJSON(&placeOrder)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.BuildResponseError(err.Error(), constants.STATUS_BAD_REQUEST, http.StatusBadRequest))
		return
	}

	token, customer, err := middleware.JWTGetTokenAndCustomer(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
		return
	}

	cartsRedis, err := h.cartService.getCart(customer.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_BAD_REQUEST, http.StatusInternalServerError))
		return
	}

	var carts []dto.CartDto
	for _, cart := range cartsRedis {
		product, err := microservices.CallGetProduct(cart.ProductId, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
			return
		}
		carts = append(carts, dto.ToCartDto(product, customer.ID, cart.Quantity))
	}

	if len(carts) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			response.BuildResponseError(constants.DB_NOT_FOUND, constants.STATUS_NOT_FOUND, http.StatusNotFound))
		return
	}

	checkOutInfo := h.cartService.checkout(carts, customer.Email)

	events := dto.PlaceOrderEvents{
		Address:       placeOrder.Address,
		PaymentMethod: placeOrder.PaymentMethod,
		CustomerId:    customer.ID,
		CustomerEmail: customer.Email,
		Items:         cartsRedis,
		CheckOut:      checkOutInfo,
	}

	jsonValue, err := json.Marshal(events)

	message := &sarama.ProducerMessage{
		Topic: constants.KAFKA_TOPIC_PLACE_ORDER,
		Value: sarama.ByteEncoder(jsonValue),
	}
	err = kafka.Produce(message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	err = h.cartService.deleteAllCart(customer.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, response.BuildResponseObject(constants.PLACE_ORDER_SUCCESS, http.StatusOK, nil))
}

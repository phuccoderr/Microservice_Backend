package cart

import (
	"cart-service/internal/constants"
	"cart-service/internal/controller/microservices"
	"cart-service/internal/dto"
	"cart-service/internal/middleware"
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
			dto.BuildResponseError(err.Error(), constants.STATUS_BADREQUEST, http.StatusBadRequest))
		return
	}

	token, err := middleware.JWTGetToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	productId := cartRequest.ProductId
	productResponse, err := microservices.CallGetProduct(productId, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			dto.BuildResponseError(err.Error(), constants.STATUS_INTERNAL_ERROR, http.StatusInternalServerError))
		return
	}

	err = h.cartService.addProductToCart(customerId, productResponse, cartRequest.Quantity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNPROCESSABLEENTITY, http.StatusUnprocessableEntity))
		return
	}

	c.JSON(http.StatusCreated, dto.BuildResponseObject(constants.CREATE_SUCCESS, http.StatusCreated, nil))

}

func (h *Handler) GetCart(c *gin.Context) {

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	carts, err := h.cartService.getCart(customerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			dto.BuildResponseError(err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, dto.BuildResponseObject(constants.GET_SUCCESS, http.StatusOK, dto.ToListCartDto(carts)))

}

func (h *Handler) DeleteCart(c *gin.Context) {
	productId := c.Param("id")

	if productId == "" {
		c.JSON(http.StatusBadRequest,
			dto.BuildResponseError("productId is required", constants.STATUS_BADREQUEST, http.StatusBadRequest))
		return
	}

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
		return
	}

	err = h.cartService.deleteCart(customerId, productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNPROCESSABLEENTITY, http.StatusUnprocessableEntity))
		return
	}

	c.JSON(http.StatusOK, dto.BuildResponseObject(constants.DELETE_SUCCESS, http.StatusOK, nil))
}

func (h *Handler) Checkout(c *gin.Context) {

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			dto.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	carts, err := h.cartService.getCart(customerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			dto.BuildResponseError(err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound))
		return
	}

	checkOutInfo := h.cartService.checkOut(carts)
	c.JSON(http.StatusOK, dto.BuildResponseObject(constants.CHECKOUT_SUCCESS, http.StatusOK, checkOutInfo))

}

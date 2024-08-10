package cart

import (
	"cart-service/internal/constants"
	"cart-service/internal/controller/microservices"
	"cart-service/internal/middleware"
	"cart-service/internal/models"
	"cart-service/internal/request"
	"cart-service/internal/response"
	"fmt"
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
	var cartRequest *request.CartRequest
	err := c.ShouldBindJSON(&cartRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.BuildResponseError(err.Error(), constants.STATUS_BADREQUEST, http.StatusBadRequest))
		return
	}

	token, err := middleware.JWTGetToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	productId := cartRequest.ProductId
	url := fmt.Sprintf("http://localhost:9140/api/v1/products/%s", productId)
	productBody := &response.ProductResponse{}

	productResponse, err := microservices.CallGetProduct(url, token, productBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			response.BuildResponseError(err.Error(), constants.STATUS_BADREQUEST, http.StatusInternalServerError))
		return
	}

	err = h.cartService.addProductToCart(customerId, productResponse, cartRequest.Quantity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			response.BuildResponseError(err.Error(), constants.STATUS_UNPROCESSABLEENTITY, http.StatusUnprocessableEntity))
		return
	}

	c.JSON(http.StatusCreated, response.BuildResponseObject(constants.CREATE_SUCCESS, http.StatusCreated, nil))

}

func (h *Handler) GetCart(c *gin.Context) {

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
	}

	cart, err := h.cartService.getCart(customerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			response.BuildResponseError(err.Error(), constants.STATUS_NOT_FOUND, http.StatusNotFound))
		return
	}

	c.JSON(http.StatusOK, response.BuildResponseObject(constants.GET_SUCCESS, http.StatusOK, models.ToCartDto(cart)))

}

func (h *Handler) DeleteCart(c *gin.Context) {
	productId := c.Param("id")

	if productId == "" {
		c.JSON(http.StatusBadRequest,
			response.BuildResponseError("productId is required", constants.STATUS_BADREQUEST, http.StatusBadRequest))
		return
	}

	customerId, err := middleware.JWTGetCustomerID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			response.BuildResponseError(err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized))
		return
	}

	err = h.cartService.deleteCart(customerId, productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			response.BuildResponseError(err.Error(), constants.STATUS_UNPROCESSABLEENTITY, http.StatusUnprocessableEntity))
		return
	}

	c.JSON(http.StatusOK, response.BuildResponseObject(constants.DELETE_SUCCESS, http.StatusOK, nil))
}

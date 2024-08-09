package cart

import (
	"cart-service/internal/constants"
	"cart-service/internal/controller/microservices"
	"cart-service/internal/middleware"
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

	token, _ := middleware.JWTGetToken(c)
	decodeToken, err := middleware.JWTDecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			response.BuildResponseError(err.Error(), constants.STATUS_BADREQUEST, http.StatusBadRequest))
		return
	}
	customerId := decodeToken.ID

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

	c.JSON(http.StatusOK, "OK")

}

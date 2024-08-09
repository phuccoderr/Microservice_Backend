package cart

import (
	"cart-service/internal/models"
	"cart-service/internal/response"
	"errors"
	"github.com/google/uuid"
	"strconv"
)

type ICartService interface {
	addProductToCart(customerId string, product *response.ProductResponse, quantity int) error
}

type cartService struct {
	Repository *CartRepository
}

func NewCartService(repository *CartRepository) ICartService {
	return &cartService{
		Repository: repository,
	}
}

func (s cartService) addProductToCart(customerId string, product *response.ProductResponse, quantity int) error {
	cart, err := s.Repository.findByCustomerAndProduct(customerId, product.Id)
	if err != nil {
		return err
	}

	updateQuantity := quantity

	if cart != nil {
		updateQuantity = cart.Quantity + updateQuantity
		if quantity > product.Stock {
			return errors.New("product stock is greater then stock:" + strconv.Itoa(product.Stock))
		}
	} else {
		cart = &models.Cart{}
		cart.Id = uuid.New().String()
		cart.CustomerId = customerId
		cart.ProductId = product.Id
	}

	cart.Quantity = updateQuantity
	s.Repository.Db.Save(cart)
	return nil
}

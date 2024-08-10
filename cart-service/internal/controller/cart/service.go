package cart

import (
	"cart-service/internal/constants"
	"cart-service/internal/models"
	"cart-service/internal/response"
	"errors"
	"github.com/google/uuid"
	"strconv"
)

type ICartService interface {
	addProductToCart(customerId string, product *response.ProductResponse, quantity int) error
	getCart(customerId string) (*models.Cart, error)
	deleteCart(customerId string, productId string) error
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

func (s cartService) getCart(customerId string) (*models.Cart, error) {
	customer, err := s.Repository.findByCustomer(customerId)
	if err != nil || customer == nil {
		return nil, errors.New(constants.DB_NOT_FOUND)
	}

	return customer, nil
}

func (s cartService) deleteCart(customerId string, productId string) error {
	err := s.Repository.deleteByCustomerAndProduct(customerId, productId)
	if err != nil {
		return err
	}
	return nil
}

package cart

import (
	"cart-service/internal/constants"
	"cart-service/internal/dto"
	"cart-service/internal/models"
	"errors"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type ICartService interface {
	addProductToCart(customerId string, product *dto.ProductResponse, quantity int) error
	getCart(customerId string) ([]models.Cart, error)
	deleteCart(customerId string, productId string) error
	checkOut(carts []models.Cart) *dto.Checkout
}

type cartService struct {
	Repository *CartRepository
}

func NewCartService(repository *CartRepository) ICartService {
	return &cartService{
		Repository: repository,
	}
}

func (s cartService) addProductToCart(customerId string, product *dto.ProductResponse, quantity int) error {
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
		cart.Price = product.Price
		cart.ProductImage = product.URL
		cart.Cost = product.Cost
	}

	cart.Quantity = updateQuantity
	cart.Total = cart.Price * float64(cart.Quantity)
	s.Repository.Db.Save(cart)
	return nil
}

func (s cartService) getCart(customerId string) ([]models.Cart, error) {
	carts, err := s.Repository.findByCustomer(customerId)
	if err != nil || carts == nil {
		return nil, errors.New(constants.DB_NOT_FOUND)
	}

	return carts, nil
}

func (s cartService) deleteCart(customerId string, productId string) error {
	err := s.Repository.deleteByCustomerAndProduct(customerId, productId)
	if err != nil {
		return err
	}
	return nil
}

func (s cartService) checkOut(carts []models.Cart) *dto.Checkout {
	checkout := &dto.Checkout{}
	for _, item := range carts {
		checkout.ProductTotal += item.Total
		checkout.ProductCost += item.Cost
	}

	checkout.DeliverDays = time.Now().AddDate(0, 0, 7)
	checkout.ShippingCost = 30.000
	checkout.ProductTotal += checkout.ShippingCost

	return checkout
}

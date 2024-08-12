package cart

import (
	"cart-service/internal/cache"
	"cart-service/internal/constants"
	"cart-service/internal/dto"
	"errors"
	"time"
)

type ICartService interface {
	addProductToCart(customerId string, productId string, quantity int64) (int64, error)
	getCart(customerId string) (map[string]string, error)
	deleteCart(customerId string, productId string) error
	checkout(carts []dto.CartDto, email string) *dto.CheckoutDto
	deleteAllCart(customerId string) error
	checkProductInCart(customerId string, productId string) (bool, error)
}

type cartService struct {
}

func NewCartService() ICartService {
	return &cartService{}
}
func (s cartService) addProductToCart(customerId string, productId string, quantity int64) (int64, error) {

	result, err := cache.Rdb.HIncrBy(cache.Ctx, cache.CartKey(customerId), productId, quantity).Result()
	if err != nil {
		return 0, err
	}

	if result < 0 {
		cache.Rdb.HDel(cache.Ctx, cache.CartKey(customerId), productId)
		return 0, errors.New("quantity cannot less than 0!")
	}
	return result, nil
}

func (s cartService) getCart(customerId string) (map[string]string, error) {
	result, err := cache.Rdb.HGetAll(cache.Ctx, cache.CartKey(customerId)).Result()
	if err != nil {
		return nil, errors.New(constants.DB_NOT_FOUND)
	}

	return result, nil
}

func (s cartService) deleteCart(customerId string, productId string) error {
	err := cache.Rdb.HDel(cache.Ctx, cache.CartKey(customerId), productId).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s cartService) checkout(carts []dto.CartDto, email string) *dto.CheckoutDto {
	checkout := &dto.CheckoutDto{}
	for _, item := range carts {
		checkout.ProductTotal += item.Total
		checkout.ProductCost = item.Cost * float64(item.Quantity)
	}

	checkout.DeliverDays = time.Now().AddDate(0, 0, 7)
	checkout.ShippingCost = 30.000
	checkout.ProductTotal += checkout.ShippingCost
	checkout.CustomerEmail = email
	return checkout
}

func (s cartService) deleteAllCart(customerId string) error {
	err := cache.Rdb.Del(cache.Ctx, cache.CartKey(customerId)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s cartService) checkProductInCart(customerId string, productId string) (bool, error) {
	exists, err := cache.Rdb.HExists(cache.Ctx, cache.CartKey(customerId), productId).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}

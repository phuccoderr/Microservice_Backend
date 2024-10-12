package service

import (
	"cart-service/internal/cache"
	"cart-service/internal/constants"
	"cart-service/internal/dto"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type ICartService interface {
	AddProductToCart(customerId string, productId string, quantity int64) (int64, error)
	GetCart(customerId string) ([]dto.CartRequest, error)
	DeleteCart(customerId string, productId string) error
	Checkout(carts []dto.CartDto, sale int64) *dto.CheckoutDto
	DeleteAllCart(customerId string) error
	CheckProductInCart(customerId string, productId string) (bool, error)
}

type cartService struct {
	Rdb *redis.Client
}

func NewCartService(rdb *redis.Client) ICartService {
	return &cartService{
		Rdb: rdb,
	}
}
func (s cartService) AddProductToCart(customerId string, productId string, quantity int64) (int64, error) {

	result, err := s.Rdb.HIncrBy(context.Background(), cache.CartKey(customerId), productId, quantity).Result()
	if err != nil {
		return 0, err
	}

	if result < 0 {
		s.Rdb.HDel(context.Background(), cache.CartKey(customerId), productId)
		return 0, errors.New("quantity cannot less than 0!")
	}
	return result, nil
}

func (s cartService) GetCart(customerId string) ([]dto.CartRequest, error) {
	result, err := s.Rdb.HGetAll(context.Background(), cache.CartKey(customerId)).Result()
	if err != nil {
		return nil, errors.New(constants.DB_NOT_FOUND)
	}

	carts := make([]dto.CartRequest, 0, len(result))
	for key, value := range result {
		quantity, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("invalid quantity for product %s: %v", key, err)
		}

		cart := dto.CartRequest{
			ProductId: key,
			Quantity:  int64(quantity),
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func (s cartService) DeleteCart(customerId string, productId string) error {
	err := s.Rdb.HDel(context.Background(), cache.CartKey(customerId), productId).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s cartService) Checkout(carts []dto.CartDto, sale int64) *dto.CheckoutDto {
	checkout := &dto.CheckoutDto{}
	for _, item := range carts {
		checkout.ProductTotal += item.Total * float64(item.Quantity)
		checkout.ProductCost += item.Cost

	}

	checkout.DeliverDays = time.Now().AddDate(0, 0, 3)
	checkout.ShippingCost = 30000

	if sale > 0 {
		var discountAmount float64
		discountAmount = checkout.ProductTotal * (float64(sale) / 100)
		checkout.ProductTotal = checkout.ProductTotal - discountAmount
		checkout.ProductTotal += checkout.ShippingCost
	}
	return checkout
}

func (s cartService) DeleteAllCart(customerId string) error {
	err := s.Rdb.Del(context.Background(), cache.CartKey(customerId)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s cartService) CheckProductInCart(customerId string, productId string) (bool, error) {
	exists, err := s.Rdb.HExists(context.Background(), cache.CartKey(customerId), productId).Result()
	if err != nil {
		return false, err
	}
	return exists, nil
}

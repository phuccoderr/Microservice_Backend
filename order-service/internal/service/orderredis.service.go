package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"order-service/global"
	"order-service/internal/cache"
	"order-service/internal/dto"
	"order-service/internal/models"
	"time"
)

type IOrderRedisService interface {
	GetAllOrder(page, limit int, sort string) ([]models.Order, error)
	SetOrder(pagination dto.PaginationDTO, page, limit int, sort string) error
}

type orderRedisService struct {
	rdb *redis.Client
}

func NewOrderRedisService(rdb *redis.Client) IOrderRedisService {
	return &orderRedisService{rdb: rdb}
}

func (ors *orderRedisService) GetAllOrder(page, limit int, sort string) ([]models.Order, error) {
	var orders []models.Order
	err := ors.rdb.Get(context.Background(), cache.OrdersKey(page, limit, sort)).Scan(&orders)
	if err == redis.Nil {
		global.Logger.Info("orders redis no content")
		return nil, err
	} else if err != nil {
		global.Logger.Info("orders redis err", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (ors *orderRedisService) SetOrder(pagination dto.PaginationDTO, page, limit int, sort string) error {
	value, err := json.Marshal(pagination)
	if err != nil {
		global.Logger.Error("marshal err", zap.Error(err))
	}
	err = ors.rdb.Set(context.Background(), cache.OrdersKey(page, limit, sort), value, 10*time.Minute).Err()
	if err != nil {
		global.Logger.Error("Redis set failed", zap.Error(err))
		return err
	}
	return nil
}

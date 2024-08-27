package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"order-service/global"
	"order-service/internal/dto"
	"time"
)

type IOrderRedisService interface {
	GetAllOrder(key string) (*dto.PaginationDTO, error)
	SetOrder(data interface{}, key string) error
	Clear(key string) error
}

type orderRedisService struct {
	rdb *redis.Client
}

func NewOrderRedisService(rdb *redis.Client) IOrderRedisService {
	return &orderRedisService{rdb: rdb}
}

func (ors *orderRedisService) GetAllOrder(key string) (*dto.PaginationDTO, error) {
	var paginationOrders dto.PaginationDTO
	result, err := ors.rdb.Get(context.Background(), key).Result()
	if err == redis.Nil {
		global.Logger.Info("orders redis no content")
		return nil, err
	} else if err != nil {
		global.Logger.Info("orders redis err", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal([]byte(result), &paginationOrders)
	if err != nil {
		global.Logger.Error("unmarshal err", zap.Error(err))
		return nil, err
	}

	return &paginationOrders, nil
}

func (ors *orderRedisService) SetOrder(data interface{}, key string) error {
	value, err := json.Marshal(data)
	if err != nil {
		global.Logger.Error("marshal err", zap.Error(err))
	}
	err = ors.rdb.Set(context.Background(), key, value, 10*time.Minute).Err()
	if err != nil {
		global.Logger.Error("Redis set failed", zap.Error(err))
		return err
	}
	return nil
}

func (ors *orderRedisService) Clear(key string) error {
	ctx := context.Background()

	keys, err := ors.rdb.Keys(ctx, key).Result()
	if err != nil {
		global.Logger.Error("redis failed to find keys", zap.Error(err))
		return err
	}

	if len(keys) > 0 {
		err = ors.rdb.Del(ctx, keys...).Err()
		if err != nil {
			global.Logger.Error("Redis del failed", zap.Error(err))
			return err
		}
	}

	return nil
}

package repository

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"order-service/global"
	"order-service/internal/models"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(mdb *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db: mdb,
	}
}

func (op *OrderRepo) FindById() string {
	return "order:123"
}

func (or *OrderRepo) FindAll(page int, limit int, sort, keyword string) ([]models.Order, error) {
	var orders []models.Order
	offset := (page - 1) * limit

	orderBy := fmt.Sprintf("created_at %s", sort)

	if err := or.db.Limit(limit).Offset(offset).Order(orderBy).Find(&orders).Error; err != nil {
		global.Logger.Error("Find All error", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

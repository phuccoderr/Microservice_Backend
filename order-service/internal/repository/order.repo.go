package repository

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"order-service/global"
	"order-service/internal/models"
)

type IOrderRepo interface {
	FindAll(page int, limit int, sort, keyword string) ([]models.Order, error)
	FindById(id string) (*models.Order, error)
	FindByCustomerId(customerId string) ([]models.Order, error)
	Save(order *models.Order) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(mdb *gorm.DB) IOrderRepo {
	return &orderRepo{
		db: mdb,
	}
}

func (or *orderRepo) FindAll(page int, limit int, sort, keyword string) ([]models.Order, error) {
	var orders []models.Order
	offset := (page - 1) * limit

	orderBy := fmt.Sprintf("created_at %s", sort)

	query := or.db.Preload("OrderDetails").Limit(limit).Offset(offset).Order(orderBy)

	if keyword != "" {
		searchKeyword := fmt.Sprintf("%%%s%%", keyword)
		query = query.Where("name LIKE ? OR address LIKE ? OR phone_number LIKE ?", searchKeyword, searchKeyword, searchKeyword)
	}

	if err := query.Find(&orders).Error; err != nil {
		global.Logger.Error("Find All error", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (or *orderRepo) FindById(id string) (*models.Order, error) {
	var order models.Order

	err := or.db.Preload("OrderDetails").First(&order, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Logger.Warn("Order not found", zap.String("id", id))
			return nil, err
		}
		global.Logger.Error("Find ById error", zap.Error(err))
		return nil, err
	}

	return &order, nil
}

func (or *orderRepo) Save(order *models.Order) error {
	if err := or.db.Save(order).Error; err != nil {
		global.Logger.Error("Save error", zap.Error(err))
		return err
	}
	return nil
}

func (or *orderRepo) FindByCustomerId(customerId string) ([]models.Order, error) {
	var orders []models.Order
	err := or.db.Preload("OrderDetails").Find(&orders, "customer_id = ?", customerId).Error
	if err != nil {
		global.Logger.Error("Find ByCustomerId error", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

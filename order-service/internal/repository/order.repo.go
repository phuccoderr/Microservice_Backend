package repository

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"order-service/global"
	"order-service/internal/models"
	"time"
)

type IOrderRepo interface {
	FindAll(page int, limit int, sort, keyword string) ([]models.Order, error)
	FindById(id string) (*models.Order, error)
	FindByCustomerId(customerId string) ([]models.Order, error)
	UpdateStatus(order *models.Order) error
	FindByPeriod(startDate, endDate time.Time) ([]models.Order, error)
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

func (or *orderRepo) UpdateStatus(order *models.Order) error {
	err := or.db.Model(order).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", order.ID).Update("status", order.Status).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Logger.Info("Record not found")
			return err
		}
		global.Logger.Error("UpdateStatus error", zap.Error(err))
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

func (or *orderRepo) FindByPeriod(startDate, endDate time.Time) ([]models.Order, error) {
	var orders []models.Order

	if err := or.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

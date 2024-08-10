package cart

import (
	"cart-service/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CartRepository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{Db: db}
}

func (r CartRepository) findByCustomerAndProduct(customerId string, productId string) (*models.Cart, error) {
	cart := &models.Cart{}

	result := r.Db.Where("customer_id = ? AND product_id = ?", customerId, productId).First(&cart)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Không tìm thấy bản ghi
			return nil, nil
		}
		return nil, result.Error
	}

	return cart, nil
}

func (r CartRepository) findByCustomer(customerId string) (*models.Cart, error) {
	cart := &models.Cart{}

	result := r.Db.Where("customer_id = ?", customerId).First(cart)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Không tìm thấy bản ghi
			return nil, nil
		}
		return nil, result.Error
	}

	return cart, nil
}

func (r CartRepository) deleteByCustomerAndProduct(customerId string, productId string) error {
	cart := &models.Cart{}

	result := r.Db.Where("customer_id = ? AND product_id = ?", customerId, productId).Delete(cart)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("no cart found with customer_id: %s and product_id: %s", customerId, productId))
	}

	return nil
}

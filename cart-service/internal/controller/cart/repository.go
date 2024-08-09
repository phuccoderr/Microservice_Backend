package cart

import (
	"cart-service/internal/models"
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
		if result.Error == gorm.ErrRecordNotFound {
			// Không tìm thấy bản ghi
			return nil, nil
		}
		return nil, result.Error
	}

	return cart, nil
}

package service

import (
	"order-service/internal/models"
	"order-service/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepo
}

func NewOrderService(repo *repository.OrderRepo) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (os *OrderService) ListByPage(page, limit int, sort, keyword string) ([]models.Order, error) {

	all, err := os.repo.FindAll(page, limit, sort, keyword)
	if err != nil {
		return nil, err
	}

	return all, nil
}

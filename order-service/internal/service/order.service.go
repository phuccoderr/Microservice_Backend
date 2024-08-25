package service

import (
	"order-service/internal/models"
	"order-service/internal/repository"
	"strings"
)

type IOrderService interface {
	ListByPage(page, limit int, sort, keyword string) ([]models.Order, error)
	FindOne(id string) (*models.Order, error)
	UpdateStatus(id string, status string) error
}

type orderService struct {
	repo repository.IOrderRepo
}

func NewOrderService(repo repository.IOrderRepo) IOrderService {
	return &orderService{
		repo: repo,
	}
}

func (os *orderService) ListByPage(page, limit int, sort, keyword string) ([]models.Order, error) {

	all, err := os.repo.FindAll(page, limit, sort, keyword)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (os *orderService) FindOne(id string) (*models.Order, error) {
	order, err := os.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *orderService) UpdateStatus(id string, status string) error {
	order, err := os.repo.FindById(id)
	if err != nil {
		return err
	}

	normalizedStatus := strings.ToLower(status)

	switch normalizedStatus {
	case models.StatusPending:
		order.Status = normalizedStatus
	case models.StatusCompleted:
		order.Status = normalizedStatus
	case models.StatusCanceled:
		order.Status = normalizedStatus
	default:
		order.Status = models.StatusPending
	}
	err = os.repo.Save(order)
	if err != nil {
		return err
	}
	return nil
}

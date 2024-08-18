package service

import "order-service/internal/repository"

type OrderService struct {
	orderRepo *repository.OrderRepo
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo: repository.NewOrderRepo(),
	}
}

func (us *OrderService) GetALL() string {
	return us.orderRepo.FindById()
}

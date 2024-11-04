package service

import (
	"errors"
	"order-service/internal/dto"
	"order-service/internal/models"
	"order-service/internal/repository"
	"strings"
	"time"
)

type IOrderService interface {
	ListByPage(page, limit int, sort, keyword string) ([]models.Order, error)
	FindOne(id string) (*models.Order, error)
	UpdateStatus(id string, status string) error
	FindByCustomer(customerId string) ([]models.Order, error)
	FindLastOptionDaysOrder(day int) ([]dto.ReportItemDto, error)
	FindLastOptionMonthsOrder(months int) ([]dto.ReportItemDto, error)
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
		return errors.New("invalid status ('pending', 'complete', 'cancel')")
	}
	err = os.repo.UpdateStatus(order)
	if err != nil {
		return err
	}
	return nil
}

func (os *orderService) FindByCustomer(customerId string) ([]models.Order, error) {
	orders, err := os.repo.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (os *orderService) FindLastOptionDaysOrder(day int) ([]dto.ReportItemDto, error) {
	var reportItems []dto.ReportItemDto
	currentDate := time.Now()

	for i := 0; i < day; i++ {
		startDate := currentDate.AddDate(0, 0, -i)
		endDate := startDate.AddDate(0, 0, 1)

		orders, err := os.repo.FindByPeriod(startDate, endDate)
		if err != nil {
			return nil, err
		}

		var grossSales, netSales float64
		for _, order := range orders {
			grossSales += order.Total
			netSales += order.Total - order.ProductCost
		}

		reportItem := dto.ReportItemDto{
			GrossSales:  grossSales,
			NetSales:    netSales,
			OrdersCount: len(orders),
			Date:        startDate,
		}

		reportItems = append(reportItems, reportItem)
	}

	return reportItems, nil
}

func (os *orderService) FindLastOptionMonthsOrder(months int) ([]dto.ReportItemDto, error) {

	var reportItems []dto.ReportItemDto
	currentDate := time.Now()

	for i := 0; i < months; i++ {
		startDate := time.Date(currentDate.Year(), currentDate.Month()-time.Month(i), 1, 0, 0, 0, 0, time.Local)
		endDate := startDate.AddDate(0, 1, 0)

		// Fetch orders for the month range in UTC and convert them to Vietnam time if necessary
		orders, err := os.repo.FindByPeriod(startDate.UTC(), endDate.UTC())
		if err != nil {
			return nil, err
		}

		var grossSales, netSales float64
		for _, order := range orders {
			grossSales += order.Total
			netSales += order.Total - order.ProductCost
		}

		reportItem := dto.ReportItemDto{
			GrossSales:  grossSales,
			NetSales:    netSales,
			OrdersCount: len(orders),
			Date:        startDate,
		}
		reportItems = append(reportItems, reportItem)
	}

	return reportItems, nil
}

package dto

import "time"

type ReportItemDto struct {
	GrossSales  float64   `json:"gross_sales"`
	NetSales    float64   `json:"net_sales"`
	OrdersCount int       `json:"orders_count"`
	Date        time.Time `json:"date"`
}

package dto

type PaginationDTO struct {
	TotalItems  int         `json:"total_items"`
	TotalPages  int         `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
	StartCount  int         `json:"start_count"`
	EndCount    int         `json:"end_count"`
	Entities    interface{} `json:"entities"`
}

func BuildPaginationDto(orders []OrderDto, page, limit int) PaginationDTO {
	startCount := (page-1)*limit + 1

	return PaginationDTO{
		TotalItems:  len(orders),
		TotalPages:  (len(orders) + limit - 1) / limit,
		CurrentPage: page,
		StartCount:  startCount,
		EndCount:    startCount + len(orders) - 1,
		Entities:    orders,
	}
}

package repository

type OrderRepo struct{}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

func (op *OrderRepo) FindById() string {
	return "order:123"
}

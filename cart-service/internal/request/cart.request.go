package request

type CartRequest struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

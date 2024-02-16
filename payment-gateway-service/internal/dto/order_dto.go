package dto

type PaymentRequest struct {
	ProductId  int `json:"product_id"`
	CostumerId int `json:"costumer_id"`
	Qty        int `json:"qty"`
}

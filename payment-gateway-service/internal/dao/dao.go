package dao

import (
	"net/http"

	"github.com/timopattikawa/payment-gateway-service/internal/domain"
)

type MidtransDao interface {
	DoRequestMidtransSnap(order domain.Order, costumer domain.Costumer) (*http.Response, error)
	DoCheckingMidtransWebhook(params string, signatureKey string) error
}

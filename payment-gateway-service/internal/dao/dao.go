package dao

import (
	"net/http"

	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

type MidtransDao interface {
	DoRequestMidtransSnap(order domain.Order, costumer *pb.Costumer) (*http.Response, error)
	DoCheckingMidtransWebhook(params string, signatureKey string) error
}

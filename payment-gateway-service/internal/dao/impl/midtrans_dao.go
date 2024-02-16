package dao

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"github.com/timopattikawa/payment-gateway-service/internal/dao"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	"github.com/timopattikawa/payment-gateway-service/internal/exception"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

type MidtransDaoImpl struct {
	cfg *config.Config
}

// DoCheckingMidtransWebhook implements dao.MidtransDao.
func (m MidtransDaoImpl) DoCheckingMidtransWebhook(params string, signatureKey string) error {
	sha := sha512.New()
	sha.Write([]byte(params + m.cfg.Midtrans.ServerKey))
	result := sha.Sum(nil)

	if string(result) != signatureKey {
		return exception.BadRequest{Message: "Signature key invalid!!!"}
	}

	return nil
}

func (m MidtransDaoImpl) DoRequestMidtransSnap(order domain.Order, costumer *pb.Costumer) (*http.Response, error) {
	client := &http.Client{}

	var payloadMidtransRequest = fmt.Sprintf(`{
		"transaction_details": {
			"order_id": "%s",
			"gross_amount": %d
		},
		"credit_card":{
			"secure" : %v
		},
		"customer_details": {
			"name": "%s",
			"email": "%s",
			"phone": "08111222333"
		}
	}`, order.UUID, order.TotalAmount, true, costumer.CostumerName, costumer.CostumerEmail)

	log.Println("Payload : ", payloadMidtransRequest)

	reqToMidtrans, err := http.NewRequest("POST", "https://app.sandbox.midtrans.com/snap/v1/transactions",
		bytes.NewBuffer([]byte(payloadMidtransRequest)))

	if err != nil {
		log.Println("Fail prepare req for midtrans err: ", err.Error())
		return nil, err
	}

	authString := base64.StdEncoding.EncodeToString([]byte(m.cfg.Midtrans.ServerKey + ":"))
	log.Println(authString)
	reqToMidtrans.Header.Add("Accept", "application/json")
	reqToMidtrans.Header.Add("Content-Type", "application/json")
	reqToMidtrans.Header.Add("Authorization",
		"Basic "+authString)

	resp, err := client.Do(reqToMidtrans)
	if err != nil {
		log.Println("Fail do req for midtrans err: ", err.Error())
		return nil, err
	}

	return resp, nil
}

func NewMidtransDao(cfg *config.Config) dao.MidtransDao {
	return &MidtransDaoImpl{
		cfg: cfg,
	}
}

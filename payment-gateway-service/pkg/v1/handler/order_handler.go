package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	"github.com/timopattikawa/payment-gateway-service/internal/dto"
	exception "github.com/timopattikawa/payment-gateway-service/internal/exception"
)

type OrderHandler struct {
	usecase domain.OrderUsecase
	cfg     *config.Config
}

func NewOrderHandler(usecase domain.OrderUsecase, cfg *config.Config) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
		cfg:     cfg,
	}
}

func (o OrderHandler) OrderRoute(app *fiber.App) {
	app.Post("order/payment", o.paymentHandler)
	app.Post("order/notification/midtrans", o.webhookHandler)
}

func (o OrderHandler) paymentHandler(c *fiber.Ctx) error {
	var paymentReq = dto.PaymentRequest{}

	if err := c.BodyParser(&paymentReq); err != nil {
		return exception.BadRequest{Message: "Bad Request please check body or params"}
	}

	result, err := o.usecase.OrderPayment(c.Context(), paymentReq)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.BaseResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (o OrderHandler) webhookHandler(c *fiber.Ctx) error {

	var midTransBody map[string]string

	if err := c.BodyParser(&midTransBody); err != nil {
		log.Println("Failed to parse body request")
		return exception.BadRequest{Message: "Bad Body not string"}
	}

	o.usecase.HandlerWebHookPayment(midTransBody)

	return c.Status(fiber.StatusOK).JSON(dto.BaseResponse{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

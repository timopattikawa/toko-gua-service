package config

import (
	"github.com/gofiber/fiber/v2"
	error2 "github.com/timopattikawa/payment-gateway-service/internal/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: error2.CustomErrorHandler,
	}
}

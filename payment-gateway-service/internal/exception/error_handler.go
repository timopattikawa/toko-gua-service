package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/payment-gateway-service/internal/dto"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {

	_, badRequestError := err.(BadRequest)
	if badRequestError {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Error Bad Request Please Check Body or Params",
			Data:    err.Error(),
		})
	}

	_, notFoundError := err.(NotFoundException)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.BaseResponse{
			Status:  fiber.StatusNotFound,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(dto.BaseResponse{
		Status:  500,
		Message: "General Error",
		Data:    err.Error(),
	})
}

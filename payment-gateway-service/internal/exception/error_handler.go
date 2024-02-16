package exception

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/payment-gateway-service/internal/dto"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {

	var badRequest BadRequest
	badRequestError := errors.As(err, &badRequest)
	if badRequestError {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.BaseResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Error Bad Request Please Check Body or Params",
			Data:    err.Error(),
		})
	}

	var notFoundException NotFoundException
	notFoundError := errors.As(err, &notFoundException)
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
		Data:    "Someting wrong with our system",
	})
}

package response

import (
	"acne-scan-api/internal/model/web"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorResponse(c *fiber.Ctx, code int, message string,errors any) error{
	return c.Status(code).JSON(web.ErrorResponse{
		Code: code,
		Message: message,
		Errors: errors,
	})

}

func InternalServerError(c *fiber.Ctx, message string, err string)error{
	logrus.Error()
	return ErrorResponse(c, http.StatusInternalServerError,message,err)
}

func BadRequest(c *fiber.Ctx, message string, err error)error{
	logrus.Error()
	return ErrorResponse(c, http.StatusBadRequest,message,err)
}
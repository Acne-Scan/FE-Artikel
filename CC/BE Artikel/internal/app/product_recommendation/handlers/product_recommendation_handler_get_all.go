package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (pr *ProductRecommendationHandlerImpl) GetAll(c *fiber.Ctx)error{
data, err := pr.productRecommendationService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "category not found", err)
		}
		return response.InternalServerError(c, "failed to get all category, something happen", err.Error())
	}

	return response.StatusOk(c,http.StatusOK,"success",data)
}
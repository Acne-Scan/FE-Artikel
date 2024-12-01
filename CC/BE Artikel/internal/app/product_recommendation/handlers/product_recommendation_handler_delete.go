package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (pr *ProductRecommendationHandlerImpl) Delete(c *fiber.Ctx) error{
	// ifExist,err:=pr.productRecommendationService.GetById(id)
	// if ifExist==nil {
	// 	return response.BadRequest(c, "cannot found article to update", err)
	// }

	idparam := c.Params("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		return response.BadRequest(c, "invalid recommendation id", err)
	}

	err = pr.productRecommendationService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "recommendation not found", err)
		}
		return response.InternalServerError(c, "failed to delete recommendation, something happen", err.Error())
	}

	return response.StatusOk(c,200,"success deleted recommendation",nil)

}
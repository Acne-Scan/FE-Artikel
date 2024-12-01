package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (articleHandler *ArticleHandlerImpl) Delete(c *fiber.Ctx) error {
	idparam := c.Params("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		return response.BadRequest(c, "invalid article id", err)
	}

	ifExist,err:=articleHandler.ArticleService.GetById(id)
	if ifExist==nil {
		return response.BadRequest(c, "cannot found article to update", err)
	}

	err = articleHandler.ArticleService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "article not found", err)
		}
		return response.InternalServerError(c, "failed to get article, something happen", err.Error())
	}

	return response.StatusOk(c,200,"success deleted article",nil)
}

package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (articleHandler *ArticleHandlerImpl) GetAll(c *fiber.Ctx) error {

	data, err := articleHandler.ArticleService.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "category not found", err)
		}
		return response.InternalServerError(c, "failed to get all article, something happen", err.Error())
	}

	return response.StatusOk(c,http.StatusOK,"success",data)
}

package handlers

import (
	"acne-scan-api/internal/app/articles/service"

	"github.com/gofiber/fiber/v2"
)

type ArticleHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error


}

type ArticleHandlerImpl struct {
	ArticleService service.ArticleService
}

func NewArticleHandlerImpl(service service.ArticleService) ArticleHandler {
	return &ArticleHandlerImpl{ArticleService: service}
}

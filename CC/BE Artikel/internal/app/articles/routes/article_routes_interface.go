package routes

import (
	"acne-scan-api/internal/app/articles/handlers"

	"github.com/gofiber/fiber/v2"
)

type ArticleRoutes interface {
	ArticleRoutes(app *fiber.App)
}

type ArticleRoutesImpl struct {
	ArticleHandler handlers.ArticleHandler
}

func NewArticleRoutes(ArticleHandler handlers.ArticleHandler) ArticleRoutes {
	return &ArticleRoutesImpl{
		ArticleHandler: ArticleHandler,
	}
}
package app

import (
	"acne-scan-api/internal/app/articles"
	productrecommendation "acne-scan-api/internal/app/product_recommendation"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func InitApp(db *sql.DB,validate *validator.Validate,app *fiber.App){
	ArticleRoutes := articles.ArticleSetup(db,validate)
	ProductRecommendationRoutes:=productrecommendation.ProductRecommendationSetup(db,validate)


	ProductRecommendationRoutes.ProductRecommendationRoutes(app)
	ArticleRoutes.ArticleRoutes(app)
}
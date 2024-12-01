package routes

import (
	"acne-scan-api/internal/app/product_recommendation/handlers"

	"github.com/gofiber/fiber/v2"
)

type ProductRecommendationRoutes interface {
	ProductRecommendationRoutes(app *fiber.App)
}

type ProductRecommendationRoutesImpl struct {
	productRecommendationHandler handlers.ProductRecommendationHandler
}

func NewProductRecommendationRoutes(productRecommendationHandler handlers.ProductRecommendationHandler) ProductRecommendationRoutes {
	return &ProductRecommendationRoutesImpl{
		productRecommendationHandler: productRecommendationHandler,
	}
}
package handlers

import (
	"acne-scan-api/internal/app/product_recommendation/service"

	"github.com/gofiber/fiber/v2"
)

type ProductRecommendationHandler interface {
	Create(c *fiber.Ctx)error
	GetAll(c *fiber.Ctx)error
	Delete(c *fiber.Ctx)error
}

type ProductRecommendationHandlerImpl struct{
	productRecommendationService service.ProductRecommendationService
}

func NewProductRecommendationHandler(service service.ProductRecommendationService)ProductRecommendationHandler{
	return &ProductRecommendationHandlerImpl{productRecommendationService:service}
}
package productrecommendation

import (
	handlerPKG "acne-scan-api/internal/app/product_recommendation/handlers"
	repositoryPKG "acne-scan-api/internal/app/product_recommendation/repository"
	routesPKG "acne-scan-api/internal/app/product_recommendation/routes"
	servicePKG "acne-scan-api/internal/app/product_recommendation/service"

	"database/sql"

	"github.com/go-playground/validator"
)

func ProductRecommendationSetup(db *sql.DB, validate *validator.Validate) routesPKG.ProductRecommendationRoutes {

	productRecommendationRepository := repositoryPKG.NewProProductRecommendation(db)
	productRecommendationService := servicePKG.NewProductRecommendationService(productRecommendationRepository, validate)
	productRecommendationHanlder := handlerPKG.NewProductRecommendationHandler(productRecommendationService)
	productRecommendationRoutes := routesPKG.NewProductRecommendationRoutes(productRecommendationHanlder)

	return productRecommendationRoutes

}

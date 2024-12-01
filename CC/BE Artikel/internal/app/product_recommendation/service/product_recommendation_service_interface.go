package service

import (
	"acne-scan-api/internal/app/product_recommendation/repository"
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"

	"github.com/go-playground/validator"
)

type ProductRecommendationService interface {
	Create(pr web.ProductRecommendationRequest)error
	GetAll() ([]domain.ProductRecommendation, error)
	Delete(id int)error

}

type ProductRecommendationServiceImpl struct{
	ProductRecommendationRepository repository.ProductRecommendationRepository
	Validator *validator.Validate
}

func NewProductRecommendationService(pr repository.ProductRecommendationRepository,validate *validator.Validate) ProductRecommendationService{
	return &ProductRecommendationServiceImpl{
		ProductRecommendationRepository: pr,
		Validator: validate,
	}
}
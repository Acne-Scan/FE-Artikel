package repository

import (
	"acne-scan-api/internal/model/domain"
	"database/sql"
)

type ProductRecommendationRepository interface {
	Create(productRecommendation *domain.ProductRecommendation) error
	GetAll() ([]domain.ProductRecommendation, error)
	Delete(id int)error
}

type ProductRecommendationImpl struct {
	DB *sql.DB
}

func NewProProductRecommendation(db *sql.DB) ProductRecommendationRepository {
	return &ProductRecommendationImpl{
		DB: db,
	}
}

package web

import "time"

type ProductRecommendationRequest struct {
	RecommendationId int       `json:"recommendation_id" validate:"required"`
	Image            string    `json:"image" validate:"required"`
	Link             string    `json:"link" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
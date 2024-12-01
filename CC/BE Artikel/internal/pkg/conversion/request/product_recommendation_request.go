package conversion

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
)

func ProductRecommendationToModel(pr web.ProductRecommendationRequest) *domain.ProductRecommendation {
	return &domain.ProductRecommendation{
		RecommendationId: pr.RecommendationId,
		Image:            pr.Image,
		Link:             pr.Link,
		Description:      pr.Description,
		CreatedAt:        pr.CreatedAt,
		UpdatedAt:        pr.UpdatedAt,
	}
}

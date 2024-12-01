package repository

import "acne-scan-api/internal/model/domain"

func (pr *ProductRecommendationImpl) GetAll() ([]domain.ProductRecommendation, error) {
	result := []domain.ProductRecommendation{}

	rows, err := pr.DB.Query("select recommendation_id , image,link, description,created_at,updated_at from pruduct_recommendation")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := domain.ProductRecommendation{}
		err = rows.Scan(
			&each.RecommendationId,
			&each.Image,
			&each.Link,
			&each.Description,
			&each.CreatedAt,
			&each.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	return result, nil
}

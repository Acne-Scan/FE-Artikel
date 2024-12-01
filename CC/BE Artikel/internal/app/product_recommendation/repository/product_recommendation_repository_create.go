package repository

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (pr *ProductRecommendationImpl) Create(productRecommendation *domain.ProductRecommendation) error {
	stmt, err := pr.DB.Prepare("insert into pruduct_recommendation (recommendation_id,image,link,description,created_at,updated_at) values (?,?,?,?,?,?)")
	if err != nil {
		return fmt.Errorf("query gagal %s", err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		productRecommendation.RecommendationId,
		productRecommendation.Image,
		productRecommendation.Link,
		productRecommendation.Description,
		productRecommendation.CreatedAt,
		productRecommendation.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

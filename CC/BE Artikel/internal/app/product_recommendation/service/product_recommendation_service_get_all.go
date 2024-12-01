package service

import "acne-scan-api/internal/model/domain"

func (pr *ProductRecommendationServiceImpl) GetAll() ([]domain.ProductRecommendation, error) {
	result, err := pr.ProductRecommendationRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return result,nil
}
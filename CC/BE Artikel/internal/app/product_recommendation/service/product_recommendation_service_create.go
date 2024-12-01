package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"time"

	"fmt"
)

func (pr *ProductRecommendationServiceImpl) Create(productRecommendation web.ProductRecommendationRequest) error {
	var err error

	err = pr.Validator.Struct(productRecommendation)
	if err != nil {
		return err
	}

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	createdAt := time.Now().In(wib)
	productRecommendation.CreatedAt = createdAt
	productRecommendation.UpdatedAt = createdAt

	conv := conversion.ProductRecommendationToModel(productRecommendation)

	err = pr.ProductRecommendationRepository.Create(conv)
	if err != nil {
		return fmt.Errorf("gagal menginput product recommendation: %s",err.Error())
	}

	return nil
}

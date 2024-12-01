package service

func (pr *ProductRecommendationServiceImpl) Delete(id int) error {

	err := pr.ProductRecommendationRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
package service

import "acne-scan-api/internal/model/domain"

func (articleService *ArticleServiceImpl) GetAll() ([]domain.Article, error) {
	data, err := articleService.ArticleRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil

}

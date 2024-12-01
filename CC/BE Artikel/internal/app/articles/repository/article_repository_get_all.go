package repository

import (
	"acne-scan-api/internal/model/domain"
)

func (articleRepository *ArticleRepositoryImpl) GetAll() ([]domain.Article, error) {

	result := []domain.Article{}

	// rows, err := articleRepository.DB.Query("select article_id, image, description, created_at,updated_at from article")

	rows, err := articleRepository.DB.Query("select article_id, name, image, description,created_at,updated_at from article")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := domain.Article{}
		err = rows.Scan(
			&each.ArticleId,
			&each.Name,
			&each.Image,
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

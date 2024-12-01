package conversion

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
)

func ArticleCreateRequestToArticleModel(article web.ArticleCreateRequest) *domain.Article {
	return &domain.Article{
		ArticleId:   article.ArticleId,
		Name:        article.Name,
		Image:       article.Image,
		Description: article.Description,
	}
}

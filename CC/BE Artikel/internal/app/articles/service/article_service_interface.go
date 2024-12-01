package service

import (
	"acne-scan-api/internal/app/articles/repository"
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"

	"github.com/go-playground/validator"
)

type ArticleService interface {
	Create(article web.ArticleCreateRequest)error
	GetAll()([]domain.Article,error)
	GetById(id int)(*domain.Article,error)
	Delete(id int)(error)
	Update(description,image string, id int)(error)
}

type ArticleServiceImpl struct{
	ArticleRepository repository.ArticleRepository
	Validator *validator.Validate
}

func NewArticleService(ar repository.ArticleRepository,validate *validator.Validate) ArticleService{
	return &ArticleServiceImpl{ArticleRepository: ar,Validator: validate}
}


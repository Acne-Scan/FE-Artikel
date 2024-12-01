package repository

import (
	"acne-scan-api/internal/model/domain"
	"database/sql"
	"time"
)

type ArticleRepository interface {
	Create(article *domain.Article)error
	Update(description,image string, id int,updatedAt time.Time)error
	Delete(id int)error
	GetAll()([]domain.Article,error)
	GetById(id int)(*domain.Article,error)
}

type ArticleRepositoryImpl struct{
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB)ArticleRepository{
	return &ArticleRepositoryImpl{DB: db}
}
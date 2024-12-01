package articles

import (
	articleHandlersPkg "acne-scan-api/internal/app/articles/handlers"
	articleRepositoryPkg "acne-scan-api/internal/app/articles/repository"
	articleRoutesPkg "acne-scan-api/internal/app/articles/routes"
	articleServicePkg "acne-scan-api/internal/app/articles/service"
	"database/sql"

	"github.com/go-playground/validator"
)

func ArticleSetup(db *sql.DB, validate *validator.Validate) articleRoutesPkg.ArticleRoutes{
	articleRepository:=articleRepositoryPkg.NewArticleRepository(db)
	articleService :=articleServicePkg.NewArticleService(articleRepository,validate)
	articleHandlers:=articleHandlersPkg.NewArticleHandlerImpl(articleService)
	articleRoutes:=articleRoutesPkg.NewArticleRoutes(articleHandlers)

	return articleRoutes
}
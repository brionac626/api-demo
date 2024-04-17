package domains

import (
	"context"

	"github.com/brionac626/api-demo/models"
)

// ArticlesRepo a repository for articles operations
type ArticlesRepo interface {
	FindAllArticles(ctx context.Context, page, limit int) ([]models.Article, int, error)
	InsertNewArticle(ctx context.Context, article models.Article) error
	UpdateArticle(ctx context.Context, article models.Article) error
	DeleteArticle(ctx context.Context, id string) error
}

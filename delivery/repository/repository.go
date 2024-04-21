package repository

import (
	"context"
	"errors"

	"github.com/brionac626/api-demo/models"
)

var (
	// ErrNoArticle is returned when the target article does not exist
	ErrNoArticle = errors.New("no article")
)

// ArticlesRepo a repository for articles operations
type ArticlesRepo interface {
	FindAllArticles(ctx context.Context, author string, page, limit int64) ([]models.Article, int64, error)
	FindArticle(ctx context.Context, id string) (*models.Article, error)
	InsertNewArticle(ctx context.Context, article models.Article) error
	UpdateArticle(ctx context.Context, article models.Article) error
	DeleteArticle(ctx context.Context, id string) error
}

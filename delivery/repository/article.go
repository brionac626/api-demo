package repository

import (
	"context"

	"github.com/brionac626/api-demo/domains"
	"github.com/brionac626/api-demo/internal/config"
	"github.com/brionac626/api-demo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ domains.ArticlesRepo = (*articleRepo)(nil)

type articleRepo struct {
	mgoDB *mongo.Database
}

// NewArticleRepo create a new article repository instance
func NewArticleRepo(mongoDBClient *mongo.Client) domains.ArticlesRepo {
	return &articleRepo{mgoDB: mongoDBClient.Database(config.GetConfig().MongoDB.DB)}
}

func (ar *articleRepo) FindAllArticles(ctx context.Context, page, limit int) ([]models.Article, int, error) {
	var result []models.Article

	return result, 0, nil
}

func (ar *articleRepo) InsertNewArticle(ctx context.Context, article models.Article) error {
	return nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, article models.Article) error {
	return nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id string) error {
	return nil
}

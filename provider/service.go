package provider

import (
	"context"

	"github.com/brionac626/api-demo/delivery/http"
	"github.com/brionac626/api-demo/delivery/repository"
	"github.com/brionac626/api-demo/internal/mongodb"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideServer(h *http.ArticleHandler) *echo.Echo {
	return http.NewServer(h)
}

func ProvideBackOfficeServer() *echo.Echo {
	return http.NewBackOfficeServer()
}

func ProvideArticleHandler(repo repository.ArticlesRepo) *http.ArticleHandler {
	return http.NewArticleHandler(repo)
}

func ProvideArticleRepository(mc *mongo.Client) repository.ArticlesRepo {
	return repository.NewArticleRepo(mc)
}

func ProvideMongoClient(ctx context.Context) (*mongo.Client, error) {
	return mongodb.InitMongoDBClient(ctx)
}

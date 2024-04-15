package domains

import "context"

type ArticlesRepo interface {
	FindAllArticles(ctx context.Context)
	InerstNewArticle(ctx context.Context)
	UpdateArticle(ctx context.Context)
	DeleteArticle(ctx context.Context, id string)
}

package repository

import (
	"context"

	"github.com/brionac626/api-demo/internal/config"
	"github.com/brionac626/api-demo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ArticlesRepo = (*articleRepo)(nil)

type articleRepo struct {
	mgoDB      *mongo.Database
	collection string
}

// NewArticleRepo create a new article repository instance
func NewArticleRepo(mongoDBClient *mongo.Client) ArticlesRepo {
	return &articleRepo{
		mgoDB:      mongoDBClient.Database(config.GetConfig().MongoDB.DB),
		collection: config.GetConfig().MongoDB.Collection,
	}
}

func (ar *articleRepo) FindAllArticles(ctx context.Context, author string, page, limit int64) ([]models.Article, int64, error) {
	invalidCount := int64(-1)
	filter := bson.D{{Key: "author", Value: author}}

	counts, err := ar.mgoDB.Collection(ar.collection).CountDocuments(ctx, filter)
	if err != nil {
		return nil, invalidCount, err
	}

	cur, err := ar.mgoDB.Collection(ar.collection).Find(
		ctx,
		filter,
		options.Find().SetSkip((page-1)*limit-1).SetLimit(limit).SetSort(bson.D{{Key: "created_at", Value: -1}}),
	)
	if err != nil {
		return nil, invalidCount, err
	}

	var result []models.Article
	if err := cur.All(ctx, &result); err != nil {
		return nil, invalidCount, err
	}

	return result, counts, nil
}

func (ar *articleRepo) FindArticle(ctx context.Context, id string) (*models.Article, error) {
	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result models.Article
	if err := ar.mgoDB.Collection(ar.collection).FindOne(
		ctx,
		bson.D{{Key: "_id", Value: articleID}},
		options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}}),
	).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (ar *articleRepo) InsertNewArticle(ctx context.Context, article models.Article) error {
	_, err := ar.mgoDB.Collection(ar.collection).InsertOne(ctx, article)
	if err != nil {
		return err
	}

	return nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, article models.Article) error {
	update := bson.D{
		{Key: "title", Value: article.Title},
		{Key: "title", Value: article.Content},
		{Key: "updated_at", Value: article.UpdatedAt},
	}
	_, err := ar.mgoDB.Collection(ar.collection).UpdateByID(ctx, article.ID, update)
	if err != nil {
		return err
	}

	return nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id string) error {
	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ar.mgoDB.Collection(ar.collection).DeleteOne(ctx, bson.D{{Key: "_id", Value: articleID}})
	if err != nil {
		return err
	}

	return nil
}

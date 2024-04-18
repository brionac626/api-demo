package repository

import (
	"context"
	"log"

	"github.com/brionac626/api-demo/internal/config"
	"github.com/brionac626/api-demo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ ArticlesRepo = (*articlesRepo)(nil)

type articlesRepo struct {
	mgoDB      *mongo.Database
	collection string
}

// NewArticleRepo create a new article repository instance
func NewArticleRepo(mongoDBClient *mongo.Client) ArticlesRepo {
	return &articlesRepo{
		mgoDB:      mongoDBClient.Database(config.GetConfig().MongoDB.DB),
		collection: config.GetConfig().MongoDB.Collection,
	}
}

func (ar *articlesRepo) FindAllArticles(ctx context.Context, author string, page, limit int64) ([]models.Article, int64, error) {
	invalidCount := int64(-1)
	filter := bson.D{{Key: "author", Value: author}}

	counts, err := ar.mgoDB.Collection(ar.collection).CountDocuments(ctx, filter)
	if err != nil {
		log.Println("count err", err)
		return nil, invalidCount, err
	}

	cur, err := ar.mgoDB.Collection(ar.collection).Find(
		ctx,
		filter,
		options.Find().SetSkip((page-1)*limit-1).SetLimit(limit).SetSort(bson.D{{Key: "created_at", Value: -1}}),
	)
	if err != nil {
		log.Println("find err", err)
		return nil, invalidCount, err
	}

	var result []models.Article
	if err := cur.All(ctx, &result); err != nil {
		log.Println("all err", err)
		return nil, invalidCount, err
	}

	return result, counts, nil
}

func (ar *articlesRepo) FindArticle(ctx context.Context, id string) (*models.Article, error) {
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

func (ar *articlesRepo) InsertNewArticle(ctx context.Context, article models.Article) error {
	_, err := ar.mgoDB.Collection(ar.collection).InsertOne(ctx, article)
	if err != nil {
		return err
	}

	return nil
}

func (ar *articlesRepo) UpdateArticle(ctx context.Context, article models.Article) error {
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

func (ar *articlesRepo) DeleteArticle(ctx context.Context, id string) error {
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

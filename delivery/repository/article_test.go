package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/brionac626/api-demo/internal/config"
	"github.com/brionac626/api-demo/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

const (
	_testDB         = "articles"
	_testCollection = "articles"
)

func TestMain(m *testing.M) {
	os.Setenv(config.EnvUnitTest, "1")
	config.InjectTestConfig(
		config.Config{
			MongoDB: config.Mongodb{
				DB:         _testDB,
				Collection: _testCollection,
			},
		},
	)

	os.Exit(m.Run())
}

func Test_articleRepo_FindAllArticles(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true).ShareClient(true))
	repo := NewArticleRepo(mt.Client)

	testAuthor := "testAuthor"
	testPage := int64(1)
	testLimit := int64(10)
	testCreatedAt := primitive.NewDateTimeFromTime(time.Now().UTC())
	testUpdatedAt := testCreatedAt

	expectedArticles := []models.Article{
		{
			ID:        primitive.NewObjectID(),
			Author:    testAuthor,
			Title:     "test-article-title-1",
			CreatedAt: testCreatedAt,
			UpdatedAt: testUpdatedAt,
			Content:   "test-article-content-1",
		},
		{
			ID:        primitive.NewObjectID(),
			Author:    testAuthor,
			Title:     "test-article-title-2",
			CreatedAt: testCreatedAt,
			UpdatedAt: testUpdatedAt,
			Content:   "test-article-content-2",
		},
		{
			ID:        primitive.NewObjectID(),
			Author:    testAuthor,
			Title:     "test-article-title-3",
			CreatedAt: testCreatedAt,
			UpdatedAt: testUpdatedAt,
			Content:   "test-article-content-3",
		},
	}
	expectedCount := int64(len(expectedArticles))

	type args struct {
		ctx    context.Context
		author string
		page   int64
		limit  int64
	}
	tests := []struct {
		name      string
		ar        ArticlesRepo
		args      args
		mockSetup func()
		want      []models.Article
		want1     int64
		wantErr   bool
	}{
		{
			name: "find-3-doc",
			ar:   repo,
			args: args{
				ctx:    context.TODO(),
				author: testAuthor,
				page:   testPage,
				limit:  testLimit,
			},
			mockSetup: func() {
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch, bson.D{{Key: "n", Value: expectedCount}}))
				mt.AddMockResponses(
					mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch,
						bson.D{
							{Key: "_id", Value: expectedArticles[0].ID},
							{Key: "author", Value: expectedArticles[0].Author},
							{Key: "title", Value: expectedArticles[0].Title},
							{Key: "content", Value: expectedArticles[0].Content},
							{Key: "created_at", Value: expectedArticles[0].CreatedAt},
							{Key: "updated_at", Value: expectedArticles[0].UpdatedAt},
						},
					),
					mtest.CreateCursorResponse(1, "articles.articles", mtest.NextBatch,
						bson.D{
							{Key: "_id", Value: expectedArticles[1].ID},
							{Key: "author", Value: expectedArticles[1].Author},
							{Key: "title", Value: expectedArticles[1].Title},
							{Key: "content", Value: expectedArticles[1].Content},
							{Key: "created_at", Value: expectedArticles[1].CreatedAt},
							{Key: "updated_at", Value: expectedArticles[1].UpdatedAt},
						},
					),
					mtest.CreateCursorResponse(0, "articles.articles", mtest.NextBatch,
						bson.D{
							{Key: "_id", Value: expectedArticles[2].ID},
							{Key: "author", Value: expectedArticles[2].Author},
							{Key: "title", Value: expectedArticles[2].Title},
							{Key: "content", Value: expectedArticles[2].Content},
							{Key: "created_at", Value: expectedArticles[2].CreatedAt},
							{Key: "updated_at", Value: expectedArticles[2].UpdatedAt},
						},
					),
				)
			},
			want:    expectedArticles,
			want1:   expectedCount,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			got, got1, err := tt.ar.FindAllArticles(tt.args.ctx, tt.args.author, tt.args.page, tt.args.limit)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
			assert.EqualValues(t, tt.want1, got1)
		})
	}
}

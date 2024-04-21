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
	"go.mongodb.org/mongo-driver/mongo"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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
			name: "found-3-articles",
			ar:   repo,
			args: args{
				ctx:    ctx,
				author: testAuthor,
				page:   testPage,
				limit:  testLimit,
			},
			mockSetup: func() {
				// count result
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch, bson.D{{Key: "n", Value: expectedCount}}))
				// find result
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
		{
			name: "no-articles-found",
			ar:   repo,
			args: args{
				ctx:    ctx,
				author: testAuthor,
				page:   testPage,
				limit:  testLimit,
			},
			mockSetup: func() {
				// count result
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch, bson.D{{Key: "n", Value: 0}}))
				// find result
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch))
			},
			want:    []models.Article{},
			want1:   0,
			wantErr: false,
		},
		{
			name: "failed-invalid-skip-number",
			ar:   repo,
			args: args{
				ctx:    ctx,
				author: testAuthor,
				page:   -1,
				limit:  0,
			},
			mockSetup: func() {
				// count result
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    400,
							Message: "invalid skip number",
							Name:    "internal",
						},
					),
				)

			},
			want:    nil,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "internal-error-when-count",
			ar:   repo,
			args: args{
				ctx:    ctx,
				author: testAuthor,
				page:   testPage,
				limit:  testLimit,
			},
			mockSetup: func() {
				// count result
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    500,
							Message: "internal error",
							Name:    "internal",
						},
					),
				)
			},
			want:    nil,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "internal-error-when-find",
			ar:   repo,
			args: args{
				ctx:    ctx,
				author: testAuthor,
				page:   testPage,
				limit:  testLimit,
			},
			mockSetup: func() {
				// count result
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch, bson.D{{Key: "n", Value: 0}}))
				// find result
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    500,
							Message: "internal error",
							Name:    "internal",
						},
					),
				)
			},
			want:    nil,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "internal-error-when-decode-result",
			ar:   repo,
			args: args{
				ctx:    ctx,
				author: testAuthor,
				page:   testPage,
				limit:  testLimit,
			},
			mockSetup: func() {
				// count result
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch, bson.D{{Key: "n", Value: expectedCount}}))
				// find result
				mt.AddMockResponses(mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch, bson.D{}))
			},
			want:    nil,
			want1:   -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			got, got1, err := tt.ar.FindAllArticles(tt.args.ctx, tt.args.author, tt.args.page, tt.args.limit)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
			assert.EqualValues(t, tt.want1, got1)
		})
	}
}

func Test_articlesRepo_FindArticle(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true).ShareClient(true))
	repo := NewArticleRepo(mt.Client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testCreatedAt := primitive.NewDateTimeFromTime(time.Now().UTC())
	testUpdatedAt := testCreatedAt
	testArticleID := primitive.NewObjectID()

	expectedArticle := &models.Article{
		ID:        testArticleID,
		Author:    "testAuthor",
		Title:     "test-article-title-1",
		CreatedAt: testCreatedAt,
		UpdatedAt: testUpdatedAt,
		Content:   "test-article-content-1",
	}

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name      string
		ar        ArticlesRepo
		args      args
		mockSetup func()
		want      *models.Article
		wantErr   bool
	}{
		{
			name: "find-1-article",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID.Hex(),
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCursorResponse(1, "articles.articles", mtest.FirstBatch,
						bson.D{
							{Key: "_id", Value: expectedArticle.ID},
							{Key: "author", Value: expectedArticle.Author},
							{Key: "title", Value: expectedArticle.Title},
							{Key: "content", Value: expectedArticle.Content},
							{Key: "created_at", Value: expectedArticle.CreatedAt},
							{Key: "updated_at", Value: expectedArticle.UpdatedAt},
						},
					),
				)
			},
			want:    expectedArticle,
			wantErr: false,
		},
		{
			name: "failed-no-article",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID.Hex(),
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    404,
							Message: mongo.ErrNoDocuments.Error(),
							Name:    "no doc",
						},
					),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed-duplicate-article",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID.Hex(),
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    11000,
							Message: "duplicate key error",
							Name:    "internal",
						},
					),
				)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed-invalid-article-id",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  "invalid-article-id",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "internal-error",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID.Hex(),
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    500,
							Message: "internal error",
							Name:    "internal",
						},
					),
				)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			got, err := tt.ar.FindArticle(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_articlesRepo_InsertNewArticle(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true).ShareClient(true))
	repo := NewArticleRepo(mt.Client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testCreatedAt := primitive.NewDateTimeFromTime(time.Now().UTC())
	testUpdatedAt := testCreatedAt
	testArticle := models.Article{
		Author:    "testAuthor",
		Title:     "test-article-title-1",
		CreatedAt: testCreatedAt,
		UpdatedAt: testUpdatedAt,
		Content:   "test-article-content-1",
	}

	type args struct {
		ctx     context.Context
		article models.Article
	}
	tests := []struct {
		name      string
		ar        ArticlesRepo
		mockSetup func()
		args      args
		wantErr   bool
	}{
		{
			name: "insert-1-article",
			ar:   repo,
			args: args{
				ctx:     ctx,
				article: testArticle,
			},
			mockSetup: func() {
				mt.AddMockResponses(mtest.CreateSuccessResponse())
			},
			wantErr: false,
		},
		{
			name: "internal-error",
			ar:   repo,
			args: args{
				ctx:     ctx,
				article: testArticle,
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    500,
							Message: "internal error",
							Name:    "internal",
						},
					),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			if err := tt.ar.InsertNewArticle(tt.args.ctx, tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("articlesRepo.InsertNewArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_articlesRepo_UpdateArticle(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true).ShareClient(true))
	repo := NewArticleRepo(mt.Client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testArticleID := primitive.NewObjectID()
	testUpdatedAt := primitive.NewDateTimeFromTime(time.Now().UTC())
	testUpdateArticle := models.Article{
		ID:        testArticleID,
		Title:     "test-article-title-1",
		UpdatedAt: testUpdatedAt,
		Content:   "test-article-content-1",
	}

	type args struct {
		ctx     context.Context
		article models.Article
	}
	tests := []struct {
		name      string
		ar        ArticlesRepo
		args      args
		mockSetup func()
		wantErr   bool
	}{
		{
			name: "update-1-article",
			ar:   repo,
			args: args{
				ctx:     ctx,
				article: testUpdateArticle,
			},
			mockSetup: func() {
				mt.AddMockResponses(bson.D{
					{Key: "ok", Value: 1},
					{Key: "n", Value: 1},
					{Key: "nModified", Value: 1},
					{Key: "upserted", Value: bson.A{
						bson.D{
							{Key: "_id", Value: testUpdateArticle.ID},
							{Key: "title", Value: testUpdateArticle.Title},
							{Key: "content", Value: testUpdateArticle.Content},
							{Key: "updated", Value: testUpdateArticle.UpdatedAt},
						},
					}},
				})
			},
			wantErr: false,
		},
		{
			name: "no-article-updated",
			ar:   repo,
			args: args{
				ctx:     ctx,
				article: testUpdateArticle,
			},

			mockSetup: func() {
				mt.AddMockResponses(bson.D{
					{Key: "ok", Value: 1},
					{Key: "n", Value: 0},
					{Key: "nModified", Value: 0},
				})
			},
			wantErr: false,
		},
		{
			name: "internal-error",
			ar:   repo,
			args: args{
				ctx:     ctx,
				article: testUpdateArticle,
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    500,
							Message: "internal error",
							Name:    "internal",
						},
					),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			if err := tt.ar.UpdateArticle(tt.args.ctx, tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("articlesRepo.UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_articlesRepo_DeleteArticle(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true).ShareClient(true))
	repo := NewArticleRepo(mt.Client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testArticleID := primitive.NewObjectID().Hex()

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name      string
		ar        ArticlesRepo
		args      args
		mockSetup func()
		wantErr   bool
	}{
		{
			name: "delete-1-article",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID,
			},
			mockSetup: func() {
				mt.AddMockResponses(bson.D{
					{Key: "ok", Value: 1},
					{Key: "n", Value: 1},
				})
			},
			wantErr: false,
		},
		{
			name: "no-article-deleted",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID,
			},
			mockSetup: func() {
				mt.AddMockResponses(bson.D{
					{Key: "ok", Value: 1},
					{Key: "n", Value: 0},
				})
			},
			wantErr: false,
		},
		{
			name: "failed-invalid-article-id",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  "invalid-article-id",
			},
			wantErr: true,
		},
		{
			name: "internal-error",
			ar:   repo,
			args: args{
				ctx: ctx,
				id:  testArticleID,
			},
			mockSetup: func() {
				mt.AddMockResponses(
					mtest.CreateCommandErrorResponse(
						mtest.CommandError{
							Code:    500,
							Message: "internal error",
							Name:    "internal",
						},
					),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			if err := tt.ar.DeleteArticle(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("articlesRepo.DeleteArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

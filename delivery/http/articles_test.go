package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brionac626/api-demo/delivery/repository"
	"github.com/brionac626/api-demo/delivery/repository/mocks"
	"github.com/brionac626/api-demo/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestArticleHandler_getArticles(t *testing.T) {
	repoMock := mocks.NewArticlesRepoMock(t)
	h := NewArticleHandler(repoMock)

	e := echo.New()
	gp := e.Group("/public")
	gp.GET("/articles/:author", h.getArticles)

	testArticleID := "123"
	testPage := int64(1)
	testLimit := int64(10)
	testAuthor := "testAuthor"
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

	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}

	tests := []struct {
		name               string
		args               args
		mockSetup          func()
		expectedStatusCode int
		expectedResp       any
		wantErr            bool
	}{
		{
			name: "find-1-article",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?id=%s&page=%d&limit=%d", testAuthor, testArticleID, testPage, testLimit), nil),
				rec: httptest.NewRecorder(),
			},
			mockSetup: func() {
				repoMock.EXPECT().FindArticle(context.Background(), testArticleID).Return(&expectedArticles[0], nil).Once()
			},
			expectedStatusCode: http.StatusOK,
			expectedResp:       models.GetArticlesResp{Articles: expectedArticles[:1], Total: 1},
			wantErr:            false,
		},
		{
			name: "failed-bad-request",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?id=%s&page=%s&limit=%s", testAuthor, testArticleID, "abc", "def"), nil),
				rec: httptest.NewRecorder(),
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResp:       models.ErrorResp{Message: "failed to bind request"},
			wantErr:            true,
		},
		{
			name: "failed-find-1-article-no-article",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?id=%s&page=%d&limit=%d", testAuthor, testArticleID, testPage, testLimit), nil),
				rec: httptest.NewRecorder(),
			},
			mockSetup: func() {
				repoMock.EXPECT().FindArticle(context.Background(), testArticleID).Return(nil, repository.ErrNoArticle).Once()
			},
			expectedStatusCode: http.StatusNotFound,
			expectedResp:       models.ErrorResp{Message: "can't find the article"},
			wantErr:            true,
		},
		{
			name: "failed-find-1-article-internal-error",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?id=%s&page=%d&limit=%d", testAuthor, testArticleID, testPage, testLimit), nil),
				rec: httptest.NewRecorder(),
			},
			mockSetup: func() {
				repoMock.EXPECT().FindArticle(context.Background(), testArticleID).Return(nil, errors.New("mongodb internal error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResp:       models.ErrorResp{Message: "find article failed"},
			wantErr:            true,
		},
		{
			name: "find-articles",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?page=%d&limit=%d", testAuthor, testPage, testLimit), nil),
				rec: httptest.NewRecorder(),
			},
			mockSetup: func() {
				repoMock.EXPECT().FindAllArticles(context.Background(), testAuthor, testPage, testLimit).Return(expectedArticles, int64(len(expectedArticles)), nil).Once()
			},
			expectedStatusCode: http.StatusOK,
			expectedResp:       models.GetArticlesResp{Articles: expectedArticles, Total: int64(len(expectedArticles))},
			wantErr:            false,
		},
		{
			name: "articles-not-found",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?page=%d&limit=%d", testAuthor, testPage, testLimit), nil),
				rec: httptest.NewRecorder(),
			},
			mockSetup: func() {
				repoMock.EXPECT().FindAllArticles(context.Background(), testAuthor, testPage, testLimit).Return([]models.Article{}, 0, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
			expectedResp:       models.GetArticlesResp{Total: 0},
			wantErr:            false,
		},
		{
			name: "failed-find-articles-internal-error",
			args: args{
				req: httptest.NewRequest(http.MethodGet, fmt.Sprintf("/public/articles/%s?page=%d&limit=%d", testAuthor, testPage, testLimit), nil),
				rec: httptest.NewRecorder(),
			},
			mockSetup: func() {
				repoMock.EXPECT().FindAllArticles(context.Background(), testAuthor, testPage, testLimit).Return(nil, -1, errors.New("mongodb internal error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResp:       models.ErrorResp{Message: "find articles failed"},
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup()
			}

			e.ServeHTTP(tt.args.rec, tt.args.req)

			assert.Equal(t, tt.expectedStatusCode, tt.args.rec.Result().StatusCode)

			body, err := io.ReadAll(tt.args.rec.Body)
			assert.NoError(t, err)

			if tt.wantErr {
				var got models.ErrorResp
				err := json.Unmarshal(body, &got)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp.(models.ErrorResp), got)
			} else {
				var got models.GetArticlesResp
				err := json.Unmarshal(body, &got)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp.(models.GetArticlesResp), got)
			}

		})
	}
}

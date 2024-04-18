package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/brionac626/api-demo/domains"
	"github.com/brionac626/api-demo/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleHandler struct {
	Repo domains.ArticlesRepo
}

func (ah *ArticleHandler) getArticles(c echo.Context) error {
	// get the list of articles from mongodb
	// order by last updated time
	// pagination, 10 per page
	ctx := c.Request().Context()

	var req models.GetArticlesReq
	if err := c.Bind(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "failed to bind request"},
		)

		return err
	}

	// find only one article
	if req.ID != nil {
		article, err := ah.Repo.FindArticle(ctx, *req.ID)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(
					http.StatusNotFound,
					&models.ErrorResp{Message: "can't find the article"},
				)
			} else {
				c.JSON(
					http.StatusInternalServerError,
					&models.ErrorResp{Message: "find article failed"},
				)
			}

			return err
		}

		return c.JSON(http.StatusOK, &article)
	}

	// find all articles per page
	req.CheckPaginationValue()

	articles, total, err := ah.Repo.FindAllArticles(ctx, req.Author, req.Page, req.Limit)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			&models.ErrorResp{Message: "find articles failed"},
		)

		return err
	}

	return c.JSON(
		http.StatusOK,
		&models.GetArticlesResp{
			Articles: articles,
			Total:    total,
		},
	)
}

func (ah *ArticleHandler) createArticles(c echo.Context) error {
	// create a new article
	// store the articles into mongodb
	ctx := c.Request().Context()

	var req models.CreateArticlesReq
	if err := c.Bind(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "failed to bind request"},
		)

		return err
	}

	articleCreatedAt := time.Now().UnixNano()

	newArticle := models.Article{
		Author:    "get-it-from-context",
		Title:     req.Title,
		CreatedAt: primitive.DateTime(articleCreatedAt),
		UpdatedAt: primitive.DateTime(articleCreatedAt),
		Content:   req.Content,
	}

	if err := ah.Repo.InsertNewArticle(ctx, newArticle); err != nil {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "insert new article error"},
		)

		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (ah *ArticleHandler) modifyArticles(c echo.Context) error {
	// modify the article's content
	// update the article's properties on mongodb
	// only the author and update the user's article
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "can't get article id"},
		)

		return errors.New("can't get article id")
	}

	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "failed to parse article id"},
		)

		return err
	}

	var req models.ModifyArticlesReq
	if err := c.Bind(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "failed to bind request"},
		)

		return err
	}

	articleUpdatedAt := time.Now().UnixNano()

	article := models.Article{
		ID:        articleID,
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: primitive.DateTime(articleUpdatedAt),
	}

	if err := ah.Repo.UpdateArticle(ctx, article); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			&models.ErrorResp{Message: "failed to update article"},
		)
	}

	return c.NoContent(http.StatusNoContent)
}

func (ah *ArticleHandler) deleteArticles(c echo.Context) error {
	// delete the articles from mongodb
	// hard delete the articles
	ctx := c.Request().Context()

	articleID := c.Param("id")
	if articleID == "" {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "can't get article id"},
		)

		return errors.New("can't get article id")
	}

	if err := ah.Repo.DeleteArticle(ctx, articleID); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			&models.ErrorResp{Message: "failed to delete article"},
		)
	}

	return c.NoContent(http.StatusNoContent)
}

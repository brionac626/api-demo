package http

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/brionac626/api-demo/delivery/repository"
	"github.com/brionac626/api-demo/models"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	repo repository.ArticlesRepo
}

func NewArticleHandler(repo repository.ArticlesRepo) *ArticleHandler {
	return &ArticleHandler{repo: repo}
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
	if req.ID != "" {
		article, err := ah.repo.FindArticle(ctx, req.ID)
		if err != nil {
			if errors.Is(err, repository.ErrNoArticle) {
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

		return c.JSON(http.StatusOK, &models.GetArticlesResp{Articles: []models.Article{*article}, Total: 1})
	}

	// find all articles per page
	req.CheckPaginationValue()

	articles, total, err := ah.repo.FindAllArticles(ctx, req.Author, req.Page, req.Limit)
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
		slog.Error("bind request error", slog.String("err", err.Error()))
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "failed to bind request"},
		)

		return err
	}

	newArticle := models.Article{
		Author:  req.Author,
		Title:   req.Title,
		Content: req.Content,
	}

	newArticle.InitArticle()

	if err := ah.repo.InsertNewArticle(ctx, newArticle); err != nil {
		c.JSON(
			http.StatusInternalServerError,
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

	var req models.ModifyArticlesReq
	if err := c.Bind(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			&models.ErrorResp{Message: "failed to bind request"},
		)

		return err
	}

	article := models.Article{
		Title:   req.Title,
		Content: req.Content,
	}

	article.ConvertArticleID(req.ID)
	article.GenerateUpdatedAtTime()

	if err := ah.repo.UpdateArticle(ctx, article); err != nil {
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

	if err := ah.repo.DeleteArticle(ctx, articleID); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			&models.ErrorResp{Message: "failed to delete article"},
		)
	}

	return c.NoContent(http.StatusNoContent)
}

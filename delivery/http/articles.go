package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getArticles(c echo.Context) error {
	// get the list of articles from mongodb
	// order by last updated time
	// pagination, 10 per page

	return c.NoContent(http.StatusNoContent)
}

func createArticles(c echo.Context) error {
	// create a new article
	// store the articles into mongodb

	return c.NoContent(http.StatusNoContent)
}

func modifyArticles(c echo.Context) error {
	// modify the article's content
	// update the article's properties on mongodb

	return c.NoContent(http.StatusNoContent)
}

func deleteArticles(c echo.Context) error {
	// delete the articles from mongodb
	// hard delete the articles

	return c.NoContent(http.StatusNoContent)
}

package http

import "github.com/labstack/echo/v4"

// NewServer return a new echo server for public service
func NewServer(h ArticleHandler) *echo.Echo {
	e := echo.New()

	// TODO: add trace

	publicGroup := e.Group("/public")

	// TODO: add articles handler
	publicGroup.GET("/articles", h.getArticles)
	publicGroup.POST("/articles", h.createArticles)
	publicGroup.PUT("/articles/:id", h.modifyArticles)
	publicGroup.DELETE("/articles/:id", h.deleteArticles)

	privateGroup := e.Group("/private")

	// TODO: add configuration handlers
	privateGroup.GET("/configuration", getConfig)     // get the current configuration
	privateGroup.POST("/configuration", updateConfig) // swap the current configuration by the new configuration
	privateGroup.PUT("/configuration/callback", nil)  // a callback that trigger the service to reload the configuration

	pFeatures := privateGroup.Group("/features")

	// TODO: add otel handlers
	pFeatures.GET("/otel", nil)
	pFeatures.PUT("/otel", nil)

	// TODO: add logger handlers
	pFeatures.GET("/logger/level", nil)
	pFeatures.PUT("/logger/level", nil)

	return e
}

// NewBackOfficeServer return a new echo service for back-office service
func NewBackOfficeServer() *echo.Echo {
	e := echo.New()
	// TODO: add back office handlers
	return e
}

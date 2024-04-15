package http

import (
	"net/http"

	"github.com/brionac626/api-demo/internal/config"
	"github.com/brionac626/api-demo/models"
	"github.com/labstack/echo/v4"
)

func getConfig(c echo.Context) error {
	config := config.GetConfig()

	return c.JSON(http.StatusOK, config)
}

func updateConfig(c echo.Context) error {
	var cfg models.Config
	if err := c.Bind(&cfg); err != nil {
		return err
	}

	config.UpdateConfig(config.Config(cfg))

	config := config.GetConfig()

	return c.JSON(http.StatusOK, config)
}

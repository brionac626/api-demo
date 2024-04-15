package provider

import (
	"github.com/brionac626/api-demo/delivery/http"

	"github.com/labstack/echo/v4"
)

func ProvideServer() *echo.Echo {
	return http.NewServer()
}

func ProvideBackOfficeServer() *echo.Echo {
	return http.NewBackOfficeServer()
}

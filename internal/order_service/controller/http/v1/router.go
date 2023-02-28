package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler *echo.Echo) {
	// Option
	handler.Use(middleware.Logger())
	handler.Use(middleware.Recover())

	// Routers
	group := handler.Group("/v1")
	newOrderServiceRoutes(group)
}

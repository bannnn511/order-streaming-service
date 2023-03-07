package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"order-streaming-services/internal/orders/usecases"
)

func NewRouter(handler *echo.Echo, uc usecases.UserCase) {
	// Option
	handler.Use(middleware.Logger())
	handler.Use(middleware.Recover())

	// K8s probe
	handler.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Routers
	group := handler.Group("/v1")
	newOrderServiceRoutes(group, uc)
}

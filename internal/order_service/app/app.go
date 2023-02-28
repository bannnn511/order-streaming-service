package app

import (
	"github.com/labstack/echo/v4"
	v1 "order-streaming-services/internal/order_service/controller/http/v1"
)

func Run() {
	handler := echo.New()
	v1.NewRouter(handler)

	handler.Logger.Fatal(handler.Start(":3000"))
}

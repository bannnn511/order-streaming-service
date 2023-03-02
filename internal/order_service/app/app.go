package app

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"order-streaming-services/cmd/order_service/config"
	v1 "order-streaming-services/internal/order_service/controller/http/v1"
)

func Run() {
	handler := echo.New()
	handler.Validator = &CustomValidator{validator: validator.New()}

	v1.NewRouter(handler)

	handler.Logger.Fatal(handler.Start(":3000"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type App struct {
	Cfg       *config.Config
	KafkaConn string
}

func New(cfg *config.Config, kafkaConn string) *App {
	return &App{
		Cfg:       cfg,
		KafkaConn: kafkaConn,
	}
}

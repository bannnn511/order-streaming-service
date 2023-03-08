package app

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"order-streaming-services/cmd/orders/config"
	v1 "order-streaming-services/internal/orders/controller/http/v1"
	"order-streaming-services/internal/orders/infras/kafka"
	"order-streaming-services/internal/orders/usecases"
	"order-streaming-services/pkg/kafka"
)

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

func Run(cfg *config.Config) {
	handler := echo.New()
	handler.Validator = &CustomValidator{validator: validator.New()}

	publisher := kafka.NewPublisher([]string{cfg.Kafka.URL})
	defer publisher.Close()

	orderServiceKafka := message.NewOrderSerViceKafka(publisher)
	uc := usecases.NewUseCase(orderServiceKafka)

	v1.NewRouter(handler, uc)

	handler.Logger.Fatal(handler.Start(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)))
}

package app

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"net/http"
	"order-streaming-services/cmd/order_service/config"
	v1 "order-streaming-services/internal/orders/controller/http/v1"
	"order-streaming-services/internal/orders/usecases"
	"order-streaming-services/internal/orders/usecases/kafka"
	kafka2 "order-streaming-services/pkg/kafka"
	"order-streaming-services/pkg/kafka/producer"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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

func Run(ctx context.Context, cancel context.CancelFunc, cfg *config.Config) {
	handler := echo.New()
	handler.Validator = &CustomValidator{validator: validator.New()}
	kafkaConn, err := kafka2.NewKafkaConn(cfg.Kafka.URL, "order-service", 0)
	if err != nil {
		slog.Error("failed to init app", err)
		cancel()
		<-ctx.Done()
	}

	messagePublisher, cleanup := producer.NewPublisher(kafkaConn.Conn)
	defer cleanup()

	orderServiceKafka := kafka.NewOrderSerViceKafka(messagePublisher)
	orderServiceKafka.Configure("order-service")
	uc := usecases.NewUseCase(orderServiceKafka)

	v1.NewRouter(handler, uc)

	handler.Logger.Fatal(handler.Start(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)))
}

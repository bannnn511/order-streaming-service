package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-streaming-services/internal/order_service/domain"
	"order-streaming-services/internal/order_service/usecases/order_service"
)

type orderServiceRoutes struct {
	uc order_service.UserCase
}

func newOrderServiceRoutes(handler *echo.Group) {
	r := &orderServiceRoutes{}

	h := handler.Group("/order-service")
	h.POST("/orders", r.createOrder)
}

type createOrderRequest struct {
	OrderBean domain.OrderBean `json:"order_bean" validate:"required"`
}

/*
createOrder persists an Order to Kafka. Returns once the order is successfully written to R nodes where
R is the replication factor configured in Kafka.
*/
func (o orderServiceRoutes) createOrder(e echo.Context) error {
	request := new(createOrderRequest)
	if err := e.Bind(request); err != nil {
		return err
	}

	if err := e.Validate(request); err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, "OK")
}

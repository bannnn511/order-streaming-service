package v1

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"order-streaming-services/internal/orders/domain"
	"order-streaming-services/internal/orders/usecases"
)

type orderServiceRoutes struct {
	uc usecases.UserCase
}

func newOrderServiceRoutes(handler *echo.Group, uc usecases.UserCase) {
	r := &orderServiceRoutes{
		uc: uc,
	}

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

	if err := o.uc.PlaceOrder(context.Background(), &request.OrderBean); err != nil {
		return fmt.Errorf("http - v1 - createOrder - %w", err)
	}

	return e.JSON(http.StatusCreated, "OK")
}

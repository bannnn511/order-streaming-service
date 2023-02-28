package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type orderServiceRoutes struct {
}

func newOrderServiceRoutes(handler *echo.Group) {
	r := &orderServiceRoutes{}

	h := handler.Group("/order-service")
	h.GET("/orders", r.createOrder)
}

/*
createOrder persists an Order to Kafka. Returns once the order is successfully written to R nodes where
R is the replication factor configured in Kafka.
*/
func (o orderServiceRoutes) createOrder(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

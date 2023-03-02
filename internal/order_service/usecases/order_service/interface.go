package order_service

import (
	"context"
	"order-streaming-services/internal/order_service/domain"
)

type UserCase interface {
	PlaceOrder(ctx context.Context, bean *domain.OrderBean) error
}

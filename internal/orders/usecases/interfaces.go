package usecases

import (
	"context"
	"order-streaming-services/internal/orders/domain"
)

type (
	OrderServiceKafka interface {
		Publish(context.Context, *domain.OrderBean) error
		Configure(topic string)
	}

	UserCase interface {
		PlaceOrder(ctx context.Context, bean *domain.OrderBean) error
	}
)

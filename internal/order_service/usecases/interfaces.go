package usecases

import (
	"context"
	"order-streaming-services/internal/order_service/domain"
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

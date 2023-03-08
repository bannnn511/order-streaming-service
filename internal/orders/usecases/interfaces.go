package usecases

import (
	"context"
	"github.com/segmentio/kafka-go"
	"order-streaming-services/internal/orders/domain"
)

type (
	OrderServiceKafka interface {
		Publish(context.Context, ...kafka.Message) error
	}

	UserCase interface {
		PlaceOrder(ctx context.Context, bean *domain.OrderBean) error
	}
)

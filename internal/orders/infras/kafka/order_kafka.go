package message

import (
	"context"
	"fmt"
	kafka2 "github.com/segmentio/kafka-go"
	"order-streaming-services/internal/orders/usecases"
	"order-streaming-services/pkg/kafka"
)

type orderServiceKafka struct {
	kafka.Publisher
}

var _ usecases.OrderServiceKafka = (*orderServiceKafka)(nil)

func NewOrderSerViceKafka(p kafka.Publisher) *orderServiceKafka {
	return &orderServiceKafka{
		p,
	}
}

func (k orderServiceKafka) Publish(ctx context.Context, msgs ...kafka2.Message) error {
	if err := k.Publisher.Publish(ctx, msgs...); err != nil {
		return fmt.Errorf("Publisher.Publish %w", err)
	}

	return nil
}

package kafka

import (
	"context"
	"fmt"
	"order-streaming-services/internal/orders/domain"
	"order-streaming-services/pkg/kafka/producer"
)

type orderServiceKafka struct {
	producer.MessagePublisher
}

func NewOrderSerViceKafka(p producer.MessagePublisher) *orderServiceKafka {
	return &orderServiceKafka{
		p,
	}
}

func (k orderServiceKafka) Publish(ctx context.Context, bean *domain.OrderBean) error {
	beanBytes, err := bean.ToByte()
	if err != nil {
		return fmt.Errorf("usecase - kafka - publish, %w", err)
	}

	if err := k.MessagePublisher.Publish(ctx, []byte(bean.Id), beanBytes); err != nil {
		return err
	}

	return nil
}

func (k orderServiceKafka) Configure(topic string) {
	k.MessagePublisher.Configure(topic)
}

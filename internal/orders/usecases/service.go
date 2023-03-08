package usecases

import (
	"context"
	"fmt"
	kafka2 "github.com/segmentio/kafka-go"
	"order-streaming-services/internal/orders/domain"
)

type usecase struct {
	serviceKafka OrderServiceKafka
}

var _ UserCase = (*usecase)(nil)

func NewUseCase(serviceKafka OrderServiceKafka) *usecase {
	return &usecase{serviceKafka: serviceKafka}
}

func (u usecase) PlaceOrder(ctx context.Context, bean *domain.OrderBean) error {
	beanBytes, err := bean.ToByte()
	if err != nil {
		return fmt.Errorf("bean.ToByte() %w", err)
	}

	message := kafka2.Message{
		Key:   []byte(bean.Id),
		Value: beanBytes,
		Topic: string(TopicOrder),
	}
	return u.serviceKafka.Publish(ctx, message)
}

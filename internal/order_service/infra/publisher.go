package infra

import "order-streaming-services/pkg/kafka/producer"

type (
	orderServiceEventPublisher struct {
		pub producer.MessagePublisher
	}
)

func NewOrderServiceEventPublisher(pub producer.MessagePublisher)

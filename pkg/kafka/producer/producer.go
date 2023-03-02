package producer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type publisher struct {
	writer *kafka.Writer
	topic  string
}

var _ MessagePublisher = (*publisher)(nil)

func NewPublisher(conn, topic string) (MessagePublisher, func() error) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(conn),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &publisher{writer: writer, topic: topic}, func() error {
		if err := writer.Close(); err != nil {
			return err
		}

		return nil
	}

}

func (p publisher) Publish(ctx context.Context, key []byte, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := p.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("kafka publish error: %w", err)
	}

	return nil
}

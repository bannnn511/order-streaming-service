package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type publisher struct {
	brokers []string
	w       *kafka.Writer
}

var _ Publisher = (*publisher)(nil)

func NewPublisher(brokers []string) *publisher {
	return &publisher{
		brokers: brokers,
		w:       NewWriter(brokers),
	}
}

func (p publisher) Publish(ctx context.Context, msgs ...kafka.Message) error {
	if err := p.w.WriteMessages(ctx, msgs...); err != nil {
		return fmt.Errorf("p.w.WriteMessages %w", err)
	}

	return nil
}

func (p publisher) Close() error {
	if err := p.w.Close(); err != nil {
		return fmt.Errorf("p.w.Close %w", err)
	}

	return nil
}

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

func NewPublisher(kafkaConn *kafka.Conn) (MessagePublisher, func() error) {
	fmt.Println(kafkaConn.RemoteAddr())
	writer := &kafka.Writer{
		Addr:                   kafkaConn.RemoteAddr(),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return &publisher{writer: writer}, func() error {
		if err := writer.Close(); err != nil {
			return err
		}

		return nil
	}

}

func (p *publisher) Configure(topic string) {
	fmt.Println("config", topic)
	p.topic = topic
	p.writer.Topic = topic
}

func (p *publisher) Publish(ctx context.Context, key []byte, value []byte) error {
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

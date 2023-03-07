package producer

import "context"

type MessagePublisher interface {
	Publish(ctx context.Context, key []byte, value []byte) error
	Configure(topic string)
}

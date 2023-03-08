package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type (
	Publisher interface {
		Publish(ctx context.Context, msgs ...kafka.Message) error
		Close() error
	}

	Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

	Consumer interface {
		ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker)
	}
)

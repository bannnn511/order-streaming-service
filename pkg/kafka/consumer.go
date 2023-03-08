package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"golang.org/x/exp/slog"
	"sync"
)

type consumer struct {
	Brokers []string
	GroupID string
	Topic   string
	reader  *kafka.Reader
}

var _ Consumer = (*consumer)(nil)

func NewConsumer(kafkaURLs []string, topic, groupId string) *consumer {
	r := NewKafkaReader(kafkaURLs, topic, groupId)

	return &consumer{
		Brokers: kafkaURLs,
		GroupID: groupId,
		Topic:   topic,
		reader:  r,
	}
}

func (c consumer) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	slog.Info("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)
	wg := &sync.WaitGroup{}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(ctx, c.reader, wg, i)
	}
	wg.Wait()
}

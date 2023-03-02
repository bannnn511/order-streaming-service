package kafka

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"golang.org/x/exp/slog"
	"time"
)

const (
	retryTimes     = 5
	backOffSeconds = 2
)

type KafkaConnStr string

var ErrCannotConnectKafka = errors.New("cannot connect to Kafka")

func NewKafkaConn(kafkaURL KafkaConnStr, topic string, partition int) (*kafka.Conn, error) {
	var (
		counts    int
		kafkaConn *kafka.Conn
	)

	for {
		leader, err := kafka.DialLeader(context.Background(), "tcp", string(kafkaURL), topic, partition)
		if err != nil {
			slog.Error("failed to connect to Kafka...", err, kafkaURL)
			counts++
		} else {
			kafkaConn = leader
			break
		}

		if counts > retryTimes {
			slog.Error("failed to retry", err)
			return nil, ErrCannotConnectKafka
		}

		slog.Info("Backing off for 2 seconds...")
		time.Sleep(backOffSeconds)
		continue
	}

	slog.Info("connected to Kafka")

	return kafkaConn, nil
}

package kafka

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"golang.org/x/exp/slog"
	"net"
	"strconv"
	"time"
)

const (
	retryTimes     = 5
	backOffSeconds = 2
)

var ErrCannotConnectKafka = errors.New("cannot connect to Kafka")

type Conn struct {
	Conn      *kafka.Conn
	topic     string
	partition int
}

func NewKafkaConn(kafkaURL string, topic string, partition int) (*Conn, error) {
	var (
		counts    int
		kafkaConn *kafka.Conn
	)
	for {
		leader, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
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

	return &Conn{
		Conn:      kafkaConn,
		topic:     topic,
		partition: partition,
	}, nil
}

func (c Conn) CreateTopic(topic string, partition int) error {
	controller, err := c.Conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partition,
			ReplicationFactor: 1,
		},
	}

	if err := controllerConn.CreateTopics(topicConfigs...); err != nil {
		return err
	}
	return nil
}

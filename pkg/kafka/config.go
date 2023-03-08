package kafka

type Config struct {
	Brokers []string
	GrouID  string
}

type TopicConfig struct {
	TopicName         string
	Partitions        int
	ReplicationFactor int
}

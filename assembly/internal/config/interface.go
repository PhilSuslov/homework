package config

import "github.com/IBM/sarama"

type KafkaConfig interface {
	Brokers() []string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type OrderAssemblyProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

type OrderPaidConsumerConfig interface {
	Topic() string
	GroupID() string
	Config() *sarama.Config
}

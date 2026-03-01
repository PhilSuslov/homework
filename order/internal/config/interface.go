package config

import "github.com/IBM/sarama"

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type inventoryGRPCConfig interface {
	Address() string
}

type paymentGRPCConfig interface {
	Address() string
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
}

type orderHTTPConfig interface {
	Address() string
	Timeout() string
}

type KafkaConfig interface {
	Brokers() []string
}

type OrderPaidProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

type OrderAssembledConsumerConfig interface {
	Topic() string
	GroupID() string
	Config() *sarama.Config
}

package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type orderAssemblyProducerEnvConfig struct {
	TopicName string `env:"ORDER_ASSEMBLED_TOPIC_NAME,required"`
}

type orderAssemblyProducerConfig struct {
	raw orderAssemblyProducerEnvConfig
}

func NewOrderAssemblyProducerConfig() (*orderAssemblyProducerConfig, error) {
	var raw orderAssemblyProducerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &orderAssemblyProducerConfig{raw: raw}, nil
}

func (cfg *orderAssemblyProducerConfig) Topic() string {
	return cfg.raw.TopicName
}

func (cfg *orderAssemblyProducerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Producer.Return.Successes = true

	return config
}

package config

import (
	"os"

	"github.com/PhilSuslov/homework/assembly/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Logger                LoggerConfig
	Kafka                 KafkaConfig
	OrderAssemblyProducer OrderAssemblyProducerConfig
	OrderPaidConsumer     OrderPaidConsumerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	orderAssemblyCfg, err := env.NewOrderAssemblyProducerConfig()
	if err != nil {
		return err
	}

	orderPaidCfg, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                loggerCfg,
		Kafka:                 kafkaCfg,
		OrderAssemblyProducer: orderAssemblyCfg,
		OrderPaidConsumer:     orderPaidCfg,
	}

	return nil
}

func AppConfig() *config { return appConfig }

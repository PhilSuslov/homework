package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"

	"github.com/PhilSuslov/homework/assembly/internal/config"
	orderconsumer "github.com/PhilSuslov/homework/assembly/internal/service/consumer/order_consumer"
	orderProducer "github.com/PhilSuslov/homework/assembly/internal/service/producer/order_producer"

	kafkaConverter "github.com/PhilSuslov/homework/assembly/internal/converter/kafka"
	decode "github.com/PhilSuslov/homework/assembly/internal/converter/kafka/decode"
	"github.com/PhilSuslov/homework/assembly/internal/service"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	wrappedKafka "github.com/PhilSuslov/homework/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/PhilSuslov/homework/platform/pkg/kafka/producer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	kafkaMiddleware "github.com/PhilSuslov/homework/platform/pkg/middleware/kafka"
)

type diContainer struct {
	consumerGroup     sarama.ConsumerGroup
	orderPaidConsumer wrappedKafka.Consumer
	orderConsumerService service.ConsumerService
	orderProducerService service.AssemblyProducerService

	orderAssemblyDecoder  kafkaConverter.AssemblyDecoder
	syncProducer          sarama.SyncProducer
	orderAssemblyProducer wrappedKafka.Producer
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssemblyProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create sync producer: %s\n", err.Error()))
		}

		closer.AddNamed("Assembly Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}

	return d.syncProducer
}

func (d *diContainer) OrderAssemblyProducer() wrappedKafka.Producer {
	if d.orderAssemblyProducer == nil {
		d.orderAssemblyProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().OrderAssemblyProducer.Topic(),
			logger.Logger(),
		)
		
	}
	return d.orderAssemblyProducer
}

func (d *diContainer) OrderPaidConsumer() wrappedKafka.Consumer {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}
	return d.orderPaidConsumer
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}

		closer.AddNamed("Assembly Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroup.Close()
		})
		d.consumerGroup = consumerGroup

	}
	return d.consumerGroup
}

func (d *diContainer) OrderAssemblyDecoder() kafkaConverter.AssemblyDecoder {
	if d.orderAssemblyDecoder == nil {
		d.orderAssemblyDecoder = decode.NewAssemblyDecode()
	}
	return d.orderAssemblyDecoder
}

func (d *diContainer) OrderConsumerService() service.ConsumerService {
	if d.orderConsumerService == nil{
		d.orderConsumerService = orderconsumer.NewService(d.OrderPaidConsumer(), 
		d.OrderAssemblyDecoder(), d.OrderAssemblyProducer())
	}
	return d.orderConsumerService
}

func (d *diContainer) OrderProducerService() service.AssemblyProducerService {
	if d.orderProducerService == nil{
		d.orderProducerService = orderProducer.NewService(d.OrderAssemblyProducer())
	}
	return d.orderProducerService
}
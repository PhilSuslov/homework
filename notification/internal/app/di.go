package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	kafkaConverter "github.com/PhilSuslov/homework/notification/internal/converter/kafka"
	kafkaDecode "github.com/PhilSuslov/homework/notification/internal/converter/kafka/decode"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	kafkaMiddleware "github.com/PhilSuslov/homework/platform/pkg/middleware/kafka"

	httpClient "github.com/PhilSuslov/homework/notification/internal/client/http"
	telegramClient "github.com/PhilSuslov/homework/notification/internal/client/http/telegram"
	"github.com/PhilSuslov/homework/notification/internal/config"
	"github.com/PhilSuslov/homework/notification/internal/service"
	telegramService "github.com/PhilSuslov/homework/notification/internal/service/telegram"

	wrappedKafka "github.com/PhilSuslov/homework/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"

	"github.com/go-telegram/bot"
)

const configPath = "../../../deploy/compose/notification/.env"

type diContainer struct {
	telegramAssembledService service.TelegramAssembledService
	telegramPaidService service.TelegramPaidService
	telegramClient  httpClient.TelegramClient
	telegramBot     *bot.Bot

	consumerGroupPaid      sarama.ConsumerGroup
	consumerGroupAssembled sarama.ConsumerGroup

	orderAssembledConsumer wrappedKafka.Consumer
	orderPaidConsumer      wrappedKafka.Consumer

	orderPaidDecoder      kafkaConverter.OrderPaidDecoder
	orderAssembledDecoder kafkaConverter.OrderAssembledDecoder
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) TelegramPaidService(ctx context.Context) service.TelegramPaidService {
	if d.telegramPaidService == nil {
		d.telegramPaidService = telegramService.NewService(d.TelegramClient(ctx))
	}
	return d.telegramPaidService
}

func (d *diContainer) TelegramAssembledService(ctx context.Context) service.TelegramAssembledService {
	if d.telegramAssembledService == nil {
		d.telegramAssembledService = telegramService.NewService(d.TelegramClient(ctx))
	}
	return d.telegramAssembledService
}


func (d *diContainer) TelegramClient(ctx context.Context) httpClient.TelegramClient {
	if d.telegramClient == nil {
		d.telegramClient = telegramClient.NewClient(d.TelegramBot(ctx))
	}
	return d.telegramClient
}

func (d *diContainer) TelegramBot(ctx context.Context) *bot.Bot {

	if d.telegramBot == nil {
		b, err := bot.New(config.AppConfig().Telegram.Token())
		if err != nil {
			panic(fmt.Sprintf("failed to create telegram bot: %s\n", err.Error()))
		}

		d.telegramBot = b
	}

	return d.telegramBot
}

func (d *diContainer) OrderAssembledConsumer() wrappedKafka.Consumer {
	if d.orderAssembledConsumer == nil {
		d.orderAssembledConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroupAssembled(),
			[]string{
				config.AppConfig().OrderAssembledConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}

	return d.orderAssembledConsumer
}

func (d *diContainer) OrderPaidConsumer() wrappedKafka.Consumer {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroupPaid(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}

	return d.orderPaidConsumer
}

func (d *diContainer) ConsumerGroupAssembled() sarama.ConsumerGroup {
	if d.consumerGroupAssembled == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssembledConsumer.GroupID(),
			config.AppConfig().OrderAssembledConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroupAssembled.Close()
		})

		d.consumerGroupAssembled = consumerGroup
	}
	return d.consumerGroupAssembled
}

func (d *diContainer) ConsumerGroupPaid() sarama.ConsumerGroup {
	if d.consumerGroupPaid == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroupPaid.Close()
		})

		d.consumerGroupPaid = consumerGroup
	}
	return d.consumerGroupPaid
}

func (d *diContainer) OrderPaidDecoder() kafkaConverter.OrderPaidDecoder {
	if d.orderPaidDecoder == nil {
		d.orderPaidDecoder = kafkaDecode.NewOrderPaidDecoder()
	}
	return d.orderPaidDecoder
}

func (d *diContainer) OrderAssembledDecoder() kafkaConverter.OrderAssembledDecoder {
	if d.orderAssembledDecoder == nil {
		d.orderAssembledDecoder = kafkaDecode.NewOrderAssembledDecoder()
	}
	return d.orderAssembledDecoder
}

package app

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	inventoryV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"github.com/PhilSuslov/homework/order/internal/config"
	"github.com/PhilSuslov/homework/order/internal/repository"
	"github.com/PhilSuslov/homework/order/internal/service"
	orderProducer "github.com/PhilSuslov/homework/order/internal/service/producer/order_producer"

	inventoryClient "github.com/PhilSuslov/homework/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/PhilSuslov/homework/order/internal/client/grpc/payment/v1"
	kafkaConverter "github.com/PhilSuslov/homework/order/internal/converter/kafka"
	"github.com/PhilSuslov/homework/order/internal/converter/kafka/decode"
	"github.com/PhilSuslov/homework/platform/pkg/closer"
	wrappedKafka "github.com/PhilSuslov/homework/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/PhilSuslov/homework/platform/pkg/kafka/producer"
	kafkaMiddleware "github.com/PhilSuslov/homework/platform/pkg/middleware/kafka"

	"github.com/PhilSuslov/homework/platform/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"

	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/order"

	orderService "github.com/PhilSuslov/homework/order/internal/service/order"
)

type diContainer struct {
	orderService    service.OrderService
	orderRepository repository.OrderRepository
	inventoryClient inventoryV1.InventoryServiceClient
	paymentClient   paymentV1.PaymentServiceClient

	postgresConn *pgxpool.Pool

	consumerGroup          sarama.ConsumerGroup
	orderAssembledConsumer wrappedKafka.Consumer

	orderRecordedDecoder kafkaConverter.OrderRecordedDecoder
	syncProducer         sarama.SyncProducer
	orderPaidProducer    wrappedKafka.Producer

	orderProducerService service.OrderProducerService
}

func NewDiContainer(ctx context.Context) *diContainer {
	d := &diContainer{}

	d.postgresConn = d.PostgresCfg(ctx)
	invClient, invConn, err := inventoryClient.NewInventoryClient()
	if err != nil {
		logger.Error(ctx, "failed to NewInventoryClient", zap.Error(err))
		return d
	}
	payClient, payConn, err := paymentClient.NewPaymentClient()
	if err != nil {
		logger.Error(ctx, "failed to NewPaymentClient", zap.Error(err))
		return d
	}

	d.inventoryClient = invClient
	d.paymentClient = payClient

	closer.AddNamed("Inventory gRPC conn", func(ctx context.Context) error {
		return invConn.Close()
	})
	closer.AddNamed("Payment gRPC conn", func(ctx context.Context) error {
		return payConn.Close()
	})

	return d
}

func (d *diContainer) OrderService(ctx context.Context) service.OrderService {
	if d.orderService == nil {
		d.orderService = orderService.NewOrderService(d.inventoryClient, d.paymentClient,
			d.OrderProducerService(ctx), d.OrderRepository(ctx))
	}
	return d.orderService
}

func (d *diContainer) OrderProducerService(ctx context.Context) service.OrderProducerService {
	if d.orderProducerService == nil {
		d.orderProducerService = orderProducer.NewService(d.OrderPaidProducer())
	}
	return d.orderProducerService
}

func (d *diContainer) OrderRepository(ctx context.Context) repository.OrderRepository {
	if d.orderRepository == nil {
		if d.postgresConn == nil {
			d.postgresConn = d.PostgresCfg(ctx)
		}
		d.orderRepository = orderRepo.NewOrderRepo(d.postgresConn)
	}
	return d.orderRepository
}

func (d *diContainer) PostgresCfg(ctx context.Context) *pgxpool.Pool {
	if d.postgresConn != nil {
		return d.postgresConn
	}

	conn, err := pgxpool.New(ctx, "postgres://demo:demo@localhost:5435/orders?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Проверяем доступность базы
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	// Миграции через Goose
	migrationsDir := "../migrations"
	db := stdlib.OpenDB(*conn.Config().ConnConfig)
	migrationsRunner := orderRepo.NewMigrator(db, migrationsDir)

	err = migrationsRunner.Up()
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	d.postgresConn = conn
	return d.postgresConn
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create sync producer: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}

	return d.syncProducer
}

func (d *diContainer) OrderPaidProducer() wrappedKafka.Producer {
	if d.orderPaidProducer == nil {
		d.orderPaidProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().OrderPaidProducer.Topic(),
			logger.Logger(),
		)
	}

	return d.orderPaidProducer
}

func (d *diContainer) OrderAssembledConsumer() wrappedKafka.Consumer {
	if d.orderAssembledConsumer == nil {
		d.orderAssembledConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().OrderAssembledConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}

	return d.orderAssembledConsumer
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssembledConsumer.GroupID(),
			config.AppConfig().OrderAssembledConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroup.Close()
		})

		d.consumerGroup = consumerGroup
	}
	return d.consumerGroup
}

func (d *diContainer) OrderRecordedDecoder() kafkaConverter.OrderRecordedDecoder {
	if d.orderRecordedDecoder == nil {
		d.orderRecordedDecoder = decode.NewOrderRecordedDecoder()
	}
	return d.orderRecordedDecoder
}

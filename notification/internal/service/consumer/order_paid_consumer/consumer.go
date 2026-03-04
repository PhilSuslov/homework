package orderpaidconsumer

import (
	"context"

	kafkaConverter "github.com/PhilSuslov/homework/notification/internal/converter/kafka"
	def "github.com/PhilSuslov/homework/notification/internal/service"
	"github.com/PhilSuslov/homework/platform/pkg/kafka"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.ConsumerService = (*service)(nil)

type service struct {
	orderPaidConsumer kafka.Consumer
	orderPaidDecoder  kafkaConverter.OrderPaidDecoder
	telegramPaid def.TelegramPaidService
}

func NewService(orderPaidConsumer kafka.Consumer,
	orderPaidDecoder kafkaConverter.OrderPaidDecoder, telegramPaid def.TelegramPaidService) *service {
	return &service{
		orderPaidConsumer: orderPaidConsumer,
		orderPaidDecoder:  orderPaidDecoder,
		telegramPaid: telegramPaid,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting orderRecordedConsumer service")

	err := s.orderPaidConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.recorder topic error", zap.Error(err))
		return err
	}

	return nil
}

package orderassembledconsumer

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
	orderAssembledConsumer kafka.Consumer
	orderAssembledDecoder  kafkaConverter.OrderAssembledDecoder
	telegramAssembled def.TelegramAssembledService

}

func NewService(orderAssembledConsumer kafka.Consumer,
	orderAssembledDecoder kafkaConverter.OrderAssembledDecoder, telegramAssembled def.TelegramAssembledService) *service {
	return &service{
		orderAssembledConsumer: orderAssembledConsumer,
		orderAssembledDecoder:  orderAssembledDecoder,
		telegramAssembled: telegramAssembled,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting orderRecordedConsumer service")

	err := s.orderAssembledConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.recorder topic error", zap.Error(err))
		return err
	}

	return nil
}

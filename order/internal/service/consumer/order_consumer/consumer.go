package orderconsumer

import (
	"context"

	kafkaConverter "github.com/PhilSuslov/homework/order/internal/converter/kafka"
	def "github.com/PhilSuslov/homework/order/internal/service"
	"github.com/PhilSuslov/homework/platform/pkg/kafka"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.ConsumerService = (*service)(nil)

type service struct {
	orderRecordedConsumer kafka.Consumer
	orderRecordedDecoder  kafkaConverter.OrderRecordedDecoder
}

func NewService(orderRecordedConsumer kafka.Consumer,
	orderRecordedDecoder kafkaConverter.OrderRecordedDecoder) *service {
	return &service{
		orderRecordedConsumer: orderRecordedConsumer,
		orderRecordedDecoder:  orderRecordedDecoder,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting orderRecordedConsumer service")

	err := s.orderRecordedConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.recorder topic error", zap.Error(err))
		return err
	}

	return nil
}

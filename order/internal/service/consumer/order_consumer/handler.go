package orderconsumer

import (
	"context"

	kafka "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) OrderHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.orderRecordedDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode OrderRecorded", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Processing message",
		zap.String("topic", msg.Topic),
		zap.Any("partition", msg.Partition),
		zap.Any("offset", msg.Offset),
		zap.String("order_uuid", event.Order_uuid),
		zap.String("payment_method", event.Payment_method),
		zap.String("transaction_uuid", event.Transaction_uuid),
	)

	return nil
}

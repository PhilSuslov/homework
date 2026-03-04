package orderassembledconsumer

import (
	"context"

	kafka "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) OrderHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.orderAssembledDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode OrderRecorded", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Processing message",
		zap.String("topic", msg.Topic),
		zap.Any("partition", msg.Partition),
		zap.Any("offset", msg.Offset),
		zap.String("order_uuid", event.Order_uuid),
		zap.String("user_uuid", event.User_uuid),
		zap.Int64("Build_time_sec", event.Build_time_sec),
	)

	err = s.telegramAssembled.SendAssembledNotification(ctx, event.User_uuid, event)
	if err != nil{
		logger.Error(ctx, "failed to send Telegram Paid Notification", zap.Error(err))
		return err
	}

	return nil
}

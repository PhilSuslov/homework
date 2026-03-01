package orderconsumer

import (
	"context"

	kafka "github.com/PhilSuslov/homework/platform/pkg/kafka/consumer"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) AssemblyHandler(ctx context.Context, msg kafka.Message) error {
	event, err := s.assemblyDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "failed to decode AssemblyHandler", zap.Error(err))
		return err
	}

	logger.Info(ctx, "Processing message",
		zap.String("Event_uuid", event.Event_uuid),
		zap.String("Order_uuid", event.Order_uuid),
		zap.String("User_uuid", event.User_uuid),
		zap.Int64("Build_time_sec", event.Build_time_sec),
	)

	return nil
}

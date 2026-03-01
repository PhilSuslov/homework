package orderconsumer

import (
	"context"

	kafkaConverter "github.com/PhilSuslov/homework/assembly/internal/converter/kafka"
	def "github.com/PhilSuslov/homework/assembly/internal/service"
	"github.com/PhilSuslov/homework/platform/pkg/kafka"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.ConsumerService = (*service)(nil)

type service struct {
	assemblyConsumer kafka.Consumer
	assemblyDecoder  kafkaConverter.AssemblyDecoder
}

func NewService(assemblyConsumer kafka.Consumer,
	assemblyDecoder kafkaConverter.AssemblyDecoder) *service {
	return &service{
		assemblyConsumer: assemblyConsumer,
		assemblyDecoder:  assemblyDecoder,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting AssemblyConsumer service")

	err := s.assemblyConsumer.Consume(ctx, s.AssemblyHandler)
	if err != nil {
		logger.Error(ctx, "Consume from AssemblyService topic error", zap.Error(err))
		return err
	}

	return nil
}

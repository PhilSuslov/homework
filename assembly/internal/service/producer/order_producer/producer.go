package orderproducer

import (
	"context"

	"github.com/PhilSuslov/homework/assembly/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/kafka"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	events_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/events/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)


// 
// const (
// 	brokerAddress = "localhost:9092" // Загрузить из .env
// 	topicName = "test-topic"
// )

type service struct{
	assemblyRecordedProducer  kafka.Producer
}

func NewService(assemblyRecordedProducer kafka.Producer) *service{
	return &service{assemblyRecordedProducer: assemblyRecordedProducer}
}

func (p *service)ProducerAssembledRecorded(ctx context.Context, event model.ShipAssembled) error{
	msg := &events_v1.ShipAssembled{
		EventUuid: event.Event_uuid,
		OrderUuid: event.Order_uuid,
		UserUuid: event.User_uuid,
		BuildTimeSec: event.Build_time_sec,
	}

	payload, err := proto.Marshal(msg)
	if err != nil{
		logger.Error(ctx, "failed to marshal msg", zap.Error(err))
		return err
	}

	err = p.assemblyRecordedProducer.Send(ctx, []byte(event.Event_uuid), payload)
	if err != nil{
		logger.Error(ctx, "failed to push assemblyRecordedProducer", zap.Error(err))
		return err
	}

	return nil
}
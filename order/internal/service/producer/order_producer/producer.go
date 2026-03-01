package orderproducer

import (
	"context"
	"log"

	"github.com/PhilSuslov/homework/order/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/kafka"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	events_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/events/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type service struct {
	orderRecordedProducer kafka.Producer
}

func NewService(orderRecordedProducer kafka.Producer) *service {
	return &service{
		orderRecordedProducer: orderRecordedProducer}
}

func (p *service) ProducerOrderRecorded(ctx context.Context, event model.OrderRecordedEvent) error {
	msg := &events_v1.OrderRecorded{
		EventUuid:       event.Event_uuid,
		OrderUuid:       event.Order_uuid,
		UserUuid:        event.User_uuid,
		PaymentMethod:   event.Payment_method,
		TransactionUuid: event.Transaction_uuid,
	}

	log.Println(msg)

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal OrderRecorded", zap.Error(err))
		return err
	}

	err = p.orderRecordedProducer.Send(ctx, []byte(event.Event_uuid), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish OrderRecorded", zap.Error(err))
		return err
	}
	return nil
}

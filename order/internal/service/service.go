package service

import (
	"context"

	"github.com/PhilSuslov/homework/order/internal/model"
	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request *model.CreateOrderRequest) (model.CreateOrderResponse, error)
	PayOrder(ctx context.Context,req *model.PayOrderRequest, orderUUID uuid.UUID) (model.PayOrderResponse, error)
	GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*model.OrderDto, error)
	CancelOrder(ctx context.Context, orderUUID uuid.UUID) ( error)
	NewError(ctx context.Context, err error) *model.GenericErrorStatusCode 
}

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type OrderProducerService interface {
	ProducerOrderRecorded(ctx context.Context, event model.OrderRecordedEvent) error
}
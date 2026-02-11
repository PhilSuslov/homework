package service

import (
	"context"


	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error)
	PayOrder(ctx context.Context,req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error)
	GetOrderByUUID(ctx context.Context, params orderV1.GetOrderByUUIDParams) (orderV1.GetOrderByUUIDRes, error)
	CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error)
}
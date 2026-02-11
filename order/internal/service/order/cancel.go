package order

import (
	"context"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	order, ok := s.orderService.CancelOrder(ctx, params)
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderV1.OrderStatusPAID {
		return nil, status.Error(codes.Unknown, "Conflict")
	}

	order.Status = orderV1.OrderStatusCANCELLED

	return nil, status.Error(codes.Canceled, "No content")
}

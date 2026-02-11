package order

import (
	"context"
	"log"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *OrderService) GetOrderByUUID(ctx context.Context, params orderV1.GetOrderByUUIDParams) (orderV1.GetOrderByUUIDRes, error){

	order, ok := s.orderService.GetOrderByUUID(ctx, params)
	if !ok {
		log.Printf("order.OrderUUID - %v. Order not found in map!", order.OrderUUID)
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &orderV1.OrderDto{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUuids,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          order.Status,
	}, nil
}
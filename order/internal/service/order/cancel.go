package order

import (
	"context"

	orderServiceModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoModel "github.com/PhilSuslov/homework/order/internal/repository/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/google/uuid"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID uuid.UUID) (error){
	order, ok := s.orderService.CancelOrder(ctx, orderUUID)
	if !ok {
		return status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderRepoModel.OrderStatus(orderServiceModel.OrderStatusPAID){
		return status.Error(codes.Unknown, "Conflict")
	}

	order.Status = orderRepoModel.OrderStatus(orderServiceModel.OrderStatusCANCELLED)


	return nil
}

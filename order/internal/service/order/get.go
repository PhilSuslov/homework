package order

import (
	"context"

	orderModel "github.com/PhilSuslov/homework/order/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"

	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (*orderModel.OrderDto, error) {

	order, ok := s.orderService.GetOrderByUUID(ctx, orderUUID)
	if !ok {
		logger.Error(ctx, "Order not found in map!", zap.String("orderUUID", orderUUID.String()))
		return nil, status.Error(codes.NotFound, "order not found")
	}

	ans := orderRepoConv.OrderDtoToService(order)

	return ans, nil

}

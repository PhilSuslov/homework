package order

import (
	"context"

	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID uuid.UUID) error {
	ok := s.orderService.CancelOrder(ctx, orderUUID)
	if !ok {
		logger.Error(ctx, "order not found", zap.String("orderUUID", orderUUID.String()))
		return status.Error(codes.NotFound, "order not found")
	}

	return nil
}

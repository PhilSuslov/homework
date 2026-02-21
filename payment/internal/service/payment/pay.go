package payment

import (
	"context"
	"errors"

	"github.com/PhilSuslov/homework/platform/pkg/logger"
	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *service) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	if req == nil {
		logger.Error(ctx, "request is nil")
		return nil, errors.New("request is nil")
	}

	tx := uuid.New()
	logger.Info(ctx, "Оплата прошла успешно! transaction_uuid: ", zap.String("tx", tx.String()))

	return &payment_v1.PayOrderResponse{
		TransactionUuid: tx.String(),
	}, nil
}

package order

import (
	"context"

	orderModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *OrderService) PayOrder(ctx context.Context,
	req *orderModel.PayOrderRequest, orderUUID uuid.UUID) (orderModel.PayOrderResponse, error) {

	order, ok := s.orderService.PayOrderCreate(ctx, orderRepoConv.PayOrderRequestToRepo(req), orderUUID)
	if !ok {
		logger.Error(ctx, "order not found", zap.Any("order", order))
		return orderModel.PayOrderResponse{}, status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusPAID) {
		logger.Error(ctx, "order already paid", zap.Any("order", order))
		return orderModel.PayOrderResponse{}, status.Error(codes.Canceled, "order already paid")
	}

	if order.Status == orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusCANCELLED) {
		logger.Error(ctx, "order cancelled", zap.Any("order", order))
		return orderModel.PayOrderResponse{}, status.Error(codes.Canceled, "order cancelled")
	}

	//Проверка метода оплаты
	var pm paymentV1.PaymentMethod
	switch req.PaymentMethod {
	case "CARD":
		pm = 1
	case "SBP":
		pm = 2
	case "CREDIT_CARD":
		pm = 3
	case "INVESTOR_MONEY":
		pm = 4
	default:
		pm = 0
	}

	// Вызываем PaymentService
	payResp, err := s.paymentClient.PayOrder(ctx, &paymentV1.PayOrderRequest{
		OrderUuid:     req.OrderUUID.String(),
		UserUuid:      order.UserUUID.String(),
		PaymentMethod: pm,
	})
	if err != nil {
		logger.Error(ctx, "payment error", zap.Any("payRes", payResp),
			zap.Error(err))
		return orderModel.PayOrderResponse{}, status.Errorf(codes.Internal, "payment error: %v", err)
	}
	transactionUUID := payResp.TransactionUuid
	paymentMethod := req.PaymentMethod

	transactionuuid, _ := uuid.Parse(transactionUUID)

	resp, err := s.orderService.PayOrder(ctx, order.OrderUUID, transactionuuid, string(paymentMethod))
	if resp == nil && err != nil {
		logger.Error(ctx, "Ошибка в service -> pay", zap.Any("resp", resp),
			zap.Error(err))

		return orderModel.PayOrderResponse{}, err
	}
	if resp == nil {
		logger.Error(ctx, "empty response from repository", zap.Error(err))
		return orderModel.PayOrderResponse{}, status.Error(codes.Internal, "empty response from repository")
	}
	respConv := orderRepoConv.PayOrderResponseToService(*resp)
	return respConv, err

}

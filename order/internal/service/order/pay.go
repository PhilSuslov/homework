package order

import (
	"context"

	orderModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *OrderService) PayOrder(ctx context.Context,
	req *orderModel.PayOrderRequest, orderUUID uuid.UUID) (orderModel.PayOrderResponse, error) {

	order, ok := s.orderService.PayOrderCreate(ctx, orderRepoConv.PayOrderRequestToRepo(req), orderUUID)

	if !ok {
		return orderModel.PayOrderResponse{}, status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusPAID) {
		return orderModel.PayOrderResponse{}, status.Error(codes.Canceled, "order already paid")
	}

	if order.Status == orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusCANCELLED) {
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
		return orderModel.PayOrderResponse{}, status.Errorf(codes.Internal, "payment error: %v", err)
	}
	transactionUUID := payResp.TransactionUuid
	paymentMethod := req.PaymentMethod

	transactionuuid, _ := uuid.Parse(transactionUUID)

	resp, err := s.orderService.PayOrder(order.OrderUUID, transactionuuid, string(paymentMethod))
	respConv := orderRepoConv.PayOrderResponseToService(*resp)
	return respConv, err

}

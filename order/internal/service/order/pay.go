package order

import (
	"context"

	orderV1 "github.com/PhilSuslov/homework/shared/pkg/openapi/order/v1"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *OrderService) PayOrder(ctx context.Context,
	req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {

	order, ok := s.orderService.PayOrderCreate(ctx, req, params)

	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	if order.Status == orderV1.OrderStatusPAID {
		return nil, status.Error(codes.Canceled, "order already paid")
	}

	if order.Status == orderV1.OrderStatusCANCELLED {
		return nil, status.Error(codes.Canceled, "order cancelled")
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
		return nil, status.Errorf(codes.Internal, "payment error: %v", err)
	}
	transactionUUID := payResp.TransactionUuid
	paymentMethod := req.PaymentMethod

	transactionuuid, _ := uuid.Parse(transactionUUID)

	resp, err := s.orderService.PayOrder(order.OrderUUID.String(), transactionuuid, paymentMethod)
	return resp, err

}

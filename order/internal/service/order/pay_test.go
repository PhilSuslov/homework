package order

import (
	"context"
	"testing"

	orderModel "github.com/PhilSuslov/homework/order/internal/model"
	orderRepo "github.com/PhilSuslov/homework/order/internal/repository/model"
	orderRepoConv "github.com/PhilSuslov/homework/order/internal/repository/converter"
	mockRepo "github.com/PhilSuslov/homework/order/internal/repository/mocks"
	paymentV1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type PayServiceSuite struct {
	suite.Suite
	ctx           context.Context
	serviceOrder  *mockRepo.OrderRepository
	service       *OrderService
	paymentClient *fakePaymentClient
}

func (s *PayServiceSuite) SetupTest() {
	s.ctx = context.Background()
	s.serviceOrder = new(mockRepo.OrderRepository)
	s.paymentClient = &fakePaymentClient{}
	s.service = &OrderService{
		orderService:  s.serviceOrder,
		paymentClient: s.paymentClient,
	}
}

// ------------------- Тесты -------------------
func (s *PayServiceSuite) TestPayOrder_Success() {
	orderUUID := uuid.New()
	userUUID := uuid.New()
	transactionUUID := uuid.New().String()

	req := &orderModel.PayOrderRequest{
		OrderUUID:     orderUUID,
		PaymentMethod: "CARD",
	}

	// Мок PayOrderCreate
	s.serviceOrder.On("PayOrderCreate",
		mock.Anything,  // ctx
		mock.Anything,  // *PayOrderRequest
		orderUUID,
	).Return(&orderRepo.OrderDto{
		OrderUUID: orderUUID,
		Status:    orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusPENDINGPAYMENT),
		UserUUID:  userUUID,
	}, true)

	// Мок PayOrder
	s.serviceOrder.On("PayOrder",
		orderUUID,
		mock.Anything,  // любой userUUID
		"CARD",
	).Return(&transactionUUID, nil)

	resp, err := s.service.PayOrder(s.ctx, req, orderUUID)
	s.Require().NoError(err)
	s.Require().NotNil(resp)

}

func (s *PayServiceSuite) TestPayOrder_AlreadyPaid() {
	orderUUID := uuid.New()
	req := &orderModel.PayOrderRequest{
		OrderUUID:     orderUUID,
		PaymentMethod: "CARD",
	}

	s.serviceOrder.On("PayOrderCreate", mock.Anything, mock.Anything, orderUUID).
		Return(&orderRepo.OrderDto{
			OrderUUID: orderUUID,
			Status:    orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusPAID),
			UserUUID:  uuid.New(),
		}, true)

	_, err := s.service.PayOrder(s.ctx, req, orderUUID)
	s.Require().Error(err)
}

func (s *PayServiceSuite) TestPayOrder_Cancelled() {
	orderUUID := uuid.New()
	req := &orderModel.PayOrderRequest{
		OrderUUID:     orderUUID,
		PaymentMethod: "CARD",
	}

	s.serviceOrder.On("PayOrderCreate", mock.Anything, mock.Anything, orderUUID).
		Return(&orderRepo.OrderDto{
			OrderUUID: orderUUID,
			Status:    orderRepoConv.OrderStatusToRepo(orderModel.OrderStatusCANCELLED),
			UserUUID:  uuid.New(),
		}, true)

	_, err := s.service.PayOrder(s.ctx, req, orderUUID)
	s.Require().Error(err)
}

func (s *PayServiceSuite) TestPayOrder_OrderNotFound() {
	orderUUID := uuid.New()
	req := &orderModel.PayOrderRequest{
		OrderUUID:     orderUUID,
		PaymentMethod: "CARD",
	}

	s.serviceOrder.On("PayOrderCreate", mock.Anything, mock.Anything, orderUUID).
		Return(nil, false)

	_, err := s.service.PayOrder(s.ctx, req, orderUUID)
	s.Require().Error(err)
}

// ------------------- Заглушка для PaymentClient -------------------

type fakePaymentClient struct{}

func (f *fakePaymentClient) PayOrder(
	ctx context.Context,
	req *paymentV1.PayOrderRequest,
	opts ...grpc.CallOption,
) (*paymentV1.PayOrderResponse, error) {
	return &paymentV1.PayOrderResponse{
		TransactionUuid: uuid.New().String(),
	}, nil
}

// ------------------- Запуск тестов -------------------

func TestPayServiceSuite(t *testing.T) {
	suite.Run(t, new(PayServiceSuite))
}
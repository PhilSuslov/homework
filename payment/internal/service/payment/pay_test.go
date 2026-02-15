package payment

import (
	"context"
	"testing"

	internal "github.com/PhilSuslov/homework/payment/internal/api/payment/v1"
	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type PayOrderSuite struct {
	suite.Suite
	servicePayment *service
	ctx            context.Context
}

func (s *PayOrderSuite) SetupTest() {
	s.servicePayment = NewPaymentService()
	s.ctx = context.Background()
}

// ===================== Unit тесты =====================

func (s *PayOrderSuite) TestPayOrder_Unit() {
	req := &payment_v1.PayOrderRequest{
		OrderUuid: uuid.New().String(),
		UserUuid:  uuid.New().String(),
	}

	resp, err := s.servicePayment.PayOrder(s.ctx, req)
	s.NoError(err)
	s.NotNil(resp)
	s.NotEmpty(resp.TransactionUuid)

	// Проверка, что это валидный UUID
	_, err = uuid.Parse(resp.TransactionUuid)
	s.NoError(err)
}

func (s *PayOrderSuite) TestPayOrder_DifferentUUIDs() {
	req := &payment_v1.PayOrderRequest{
		OrderUuid: uuid.New().String(),
		UserUuid:  uuid.New().String(),
	}

	resp1, _ := s.servicePayment.PayOrder(s.ctx, req)
	resp2, _ := s.servicePayment.PayOrder(s.ctx, req)

	s.NotEqual(resp1.TransactionUuid, resp2.TransactionUuid)
}

// ===================== Интеграционные тесты через gRPC =====================

func (s *PayOrderSuite) TestPayOrder_IntegrationGRPC() {
	// Поднимаем сервер на случайном порту
	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)

	grpcServer := grpc.NewServer()
	api := internal.NewAPI()
	payment_v1.RegisterPaymentServiceServer(grpcServer, api)
	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	// Создаём клиента
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	defer conn.Close()

	client := payment_v1.NewPaymentServiceClient(conn)

	req := &payment_v1.PayOrderRequest{
		OrderUuid: uuid.New().String(),
		UserUuid:  uuid.New().String(),
	}

	resp, err := client.PayOrder(s.ctx, req)
	s.NoError(err)
	s.NotEmpty(resp.TransactionUuid)

	// Проверка на корректность UUID
	_, err = uuid.Parse(resp.TransactionUuid)
	s.NoError(err)
}

// ===================== Тест ошибки =====================

func (s *PayOrderSuite) TestPayOrder_NilRequest() {
	_, err := s.servicePayment.PayOrder(s.ctx, nil)
	assert.Error(s.T(), err) // Если сервис будет проверять nil request
}

func TestPayOrderSuite(t *testing.T) {
	suite.Run(t, new(PayOrderSuite))
}
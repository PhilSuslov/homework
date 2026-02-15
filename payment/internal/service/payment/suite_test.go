package payment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	// mock "github.com/stretchr/testify/mock"
	mock "github.com/PhilSuslov/homework/payment/internal/service/mocks"
	// servicePayment "github.com/PhilSuslov/homework/payment/internal/service"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	servicePayment *mock.PayService

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.servicePayment = mock.NewPayService(s.T())

	s.service = NewPaymentService()
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

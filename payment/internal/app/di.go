package app

import (
	"context"

	api "github.com/PhilSuslov/homework/payment/internal/api/payment/v1"
	"github.com/PhilSuslov/homework/payment/internal/service"
	"github.com/PhilSuslov/homework/payment/internal/service/payment"

	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
)

type diContainer struct {
	paymentV1API   payment_v1.PaymentServiceServer
	paymentService service.PayService
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) PaymentV1API(ctx context.Context) payment_v1.PaymentServiceServer {
	if d.paymentV1API == nil {
		d.paymentV1API = api.NewAPI()
	}
	return d.paymentV1API
}

func (d *diContainer) PartService(ctx context.Context) service.PayService {
	if d.paymentService == nil {
		d.paymentService = payment.NewPaymentService()
	}
	return d.paymentService
}

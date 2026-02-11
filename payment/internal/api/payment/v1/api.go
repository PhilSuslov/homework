package v1

import (
	"github.com/PhilSuslov/homework/payment/internal/service"
	"github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
)

type api struct {
	payment_v1.UnimplementedPaymentServiceServer

	PayService service.PayService
}

func NewAPI() *api{
	return &api{}
}
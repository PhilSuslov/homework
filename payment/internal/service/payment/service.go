package payment

import (
	def "github.com/PhilSuslov/homework/payment/internal/service"
)


var _ def.PayService = (*service)(nil)

type service struct{
}

func NewPaymentService() *service {
	return &service{
	}
}
package service

import (
	"context"

	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
)

type PayService interface {
	PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error)
}

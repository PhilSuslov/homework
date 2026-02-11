package payment

import (
	"context"
	"log"

	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"github.com/google/uuid"
)

func (s *service) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	tx := uuid.New()
	log.Println("PAY ORDER is PAYMENT ")
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", tx)

	return &payment_v1.PayOrderResponse{
		TransactionUuid: tx.String(),
	}, nil
}


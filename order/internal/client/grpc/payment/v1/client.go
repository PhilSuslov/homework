package v1

import (
	"log"

	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPaymentClient() (payment_v1.PaymentServiceClient, *grpc.ClientConn, error) {
	payConn, err := grpc.NewClient(
		"dns:///localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to payment: %v", err)
	}

	paymentClient := payment_v1.NewPaymentServiceClient(payConn)
	return paymentClient, payConn, nil
}

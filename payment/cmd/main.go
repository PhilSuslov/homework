package main

import (
	"context"
	"log"
	"net"

	pb "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) PayOrder(ctx context.Context, req *pb.PayOrderRequest) (*pb.PayOrderResponse, error) {
	tx := uuid.New()
	log.Println("PAY ORDER is PAYMENT ")
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", tx)

	return &pb.PayOrderResponse{
		TransactionUuid: tx.String(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	paymentService := NewPaymentService()
	pb.RegisterPaymentServiceServer(grpcServer, paymentService)

	log.Println("💳 Payment service started on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

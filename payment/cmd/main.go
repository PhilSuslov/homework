package main

import (
	"log"
	"net"

	paymentAPI "github.com/PhilSuslov/homework/payment/internal/api/payment/v1"
	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	paymentAPI := paymentAPI.NewAPI()
	payment_v1.RegisterPaymentServiceServer(grpcServer, paymentAPI)

	log.Println("💳 Payment service started on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

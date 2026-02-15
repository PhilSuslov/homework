package internal

import (
	"log"
	"net"

	paymentAPI "github.com/PhilSuslov/homework/payment/internal/api/payment/v1"
	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartPaymentServer(port string) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	api := paymentAPI.NewAPI()
	payment_v1.RegisterPaymentServiceServer(grpcServer, api)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	return grpcServer, lis, nil
}
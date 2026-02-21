package internal

import (
	"context"
	"net"

	paymentAPI "github.com/PhilSuslov/homework/payment/internal/api/payment/v1"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	payment_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/payment/v1"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartPaymentServer(port string) (*grpc.Server, net.Listener, error) {
	ctx := context.Background()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error(ctx, "failed to connect PaymentServer in port:", zap.String("port", port),
			zap.Error(err))
		return nil, nil, err
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	api := paymentAPI.NewAPI()
	payment_v1.RegisterPaymentServiceServer(grpcServer, api)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal(ctx, "failed to connect gRPS server", zap.Error(err))
		}
	}()

	return grpcServer, lis, nil
}

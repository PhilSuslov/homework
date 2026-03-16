package v1

import (
	"log"

	auth_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
	user_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"


	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient() (auth_v1.AuthServiceClient, *grpc.ClientConn, error) {
	authConn, err := grpc.NewClient(
		"dns:///localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to AuthClient: %v", err)
	}

	authClient := auth_v1.NewAuthServiceClient(authConn)
	return authClient, authConn, nil

}

func NewUserClient() (user_v1.UserServiceClient, *grpc.ClientConn, error) {
	userConn, err := grpc.NewClient(
		"dns:///localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to UserClient: %v", err)
	}

	userClient := user_v1.NewUserServiceClient(userConn)
	return userClient, userConn, nil

}
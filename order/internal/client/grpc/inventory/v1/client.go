package v1

import (
	"log"

	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewInventoryClient() (inventory_v1.InventoryServiceClient, *grpc.ClientConn, error) {
	invConn, err := grpc.NewClient(
		"dns:///localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to inventory: %v", err)
	}

	inventoryClient := inventory_v1.NewInventoryServiceClient(invConn)
	return inventoryClient, invConn, nil

}

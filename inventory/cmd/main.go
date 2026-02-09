package main

import (
	"log"
	"net"

	pb "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
    api "github.com/PhilSuslov/homework/inventory/internal/api/inventory/v1/"


	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"

)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	inventoryService := NewInventoryService()
	pb.RegisterInventoryServiceServer(grpcServer, inventoryService)

	log.Println("📦 Inventory service started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

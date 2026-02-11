package main

import (
	"log"
	"net"

	inventoryAPI "github.com/PhilSuslov/homework/inventory/internal/api/inventory/v1"
	inventoryRepo "github.com/PhilSuslov/homework/inventory/internal/repository/part"
	inventoryService "github.com/PhilSuslov/homework/inventory/internal/service/part"
	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("Failed to closer listener: %v\n", cerr)
		}
	}()

	grpcServer := grpc.NewServer()

	// Регистрируем наш сервис
	repo := inventoryRepo.NewRepository()
	service := inventoryService.NewService(repo)
	api := inventoryAPI.NewAPI(service)

	inventory_v1.RegisterInventoryServiceServer(grpcServer, api)
	reflection.Register(grpcServer)

	log.Println("📦 Inventory service started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

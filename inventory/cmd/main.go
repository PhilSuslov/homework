package main

import (
	inV1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
)

type InventoryService struct {
	inV1.UnimplementedInventoryServiceServer
	parts map[string]*inV1.Part
}

func NewInventoryServiceClient() *InventoryService {
	return &InventoryService{
		parts: make(map[string]*inV1.Part),
	}
}

func main() {
}

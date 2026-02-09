package v1

import (
	pb "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"github.com/google/uuid"

)

type InventoryService struct {
	pb.UnimplementedInventoryServiceServer
	parts map[string]*pb.Part
}

func NewInventoryService() *InventoryService {
	service := &InventoryService{
		parts: make(map[string]*pb.Part),
	}

	id := uuid.New()
	service.parts[id.String()] = &pb.Part{
		Uuid: id.String(),
		Name: "Main Engine",
		Price: 1445.31,
		Category: pb.Category_CATEGORY_ENGINE,
		Manufacturer: &pb.Manufacturer{Country: "German"},
		Tags: []string{"main", "engine"},
	}
	
	return service
}

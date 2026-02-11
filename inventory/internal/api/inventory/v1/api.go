package v1

import (
	"github.com/PhilSuslov/homework/inventory/internal/service"
	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"


)


type api struct {
	inventory_v1.UnimplementedInventoryServiceServer
	
	inventoryService service.InventoryService
}

func NewAPI (inventoryService service.InventoryService) *api{
	return &api{
		inventoryService: inventoryService,
	}
}
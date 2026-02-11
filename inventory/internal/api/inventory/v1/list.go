package v1

import (
	"context"
	"log"

	conv "github.com/PhilSuslov/homework/inventory/internal/converter"
	// "github.com/PhilSuslov/homework/inventory/internal/model"
	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(ctx context.Context, req *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
	part, err := a.inventoryService.ListParts(ctx, conv.InventoryListPartsRequestToModel(req))
	log.Println("List part is API:", part)
	if err != nil {
		return nil, err
	}

	return &inventory_v1.ListPartsResponse{
		Parts: conv.InventoryListPartsResponseToModel(part).Parts,
	}, nil
}

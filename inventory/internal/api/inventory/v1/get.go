package v1

import (
	"context"

	conv "github.com/PhilSuslov/homework/inventory/internal/converter"
	"github.com/PhilSuslov/homework/inventory/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"go.uber.org/zap"
)

func (a *api) GetPart(ctx context.Context, req *inventory_v1.GetPartRequest) (*inventory_v1.GetPartResponse, error) {
	part, err := a.inventoryService.GetPart(ctx, conv.InventoryGetToModel(req))

	if err != nil {
		logger.Error(ctx, "failed to GetPart in api/inventory", 
		zap.Error(model.ErrNotFound))
		
		return nil, model.ErrNotFound
	}

	return &inventory_v1.GetPartResponse{
		Part: conv.InventoryGetPartResponseToProto(part.Part.Body),
	}, nil

}

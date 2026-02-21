package v1

import (
	"context"

	conv "github.com/PhilSuslov/homework/inventory/internal/converter"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
	"go.uber.org/zap"
)

func (a *api) ListParts(ctx context.Context, req *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
	part, err := a.inventoryService.ListParts(ctx, conv.InventoryListPartsRequestToModel(req))
	if err != nil {
		logger.Error(ctx, "failed to ListParts in api/inventory",
		zap.Error(err))

		return nil, err
	}

	return &inventory_v1.ListPartsResponse{
		Parts: conv.InventoryListPartsResponseToModel(part).Parts,
	}, nil
}

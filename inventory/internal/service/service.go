package service

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	// repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

type InventoryService interface {
	GetPart(ctx context.Context, req model.GetPartRequest) (model.GetPartResponse, error)
	ListParts(ctx context.Context, req model.ListPartsRequest) (model.ListPartsResponse, error)
}

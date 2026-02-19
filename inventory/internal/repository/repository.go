package repository

import (
	"context"

	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

type InventoryRepository interface {
	GetPart(ctx context.Context, info repoModel.GetPartRequest) (repoModel.GetPartResponse, error)
	ListParts(ctx context.Context, req repoModel.ListPartsRequest) (repoModel.ListPartsResponse, error)
}
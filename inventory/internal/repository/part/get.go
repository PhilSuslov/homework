package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	repoConverter "github.com/PhilSuslov/homework/inventory/internal/repository/converter"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

func (r *Repository) GetPart(ctx context.Context, info model.GetPartRequest) (model.GetPartResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.Parts[info.Uuid]
	if !ok {
		return model.GetPartResponse{}, repoModel.ErrNotFound
	}

	return model.GetPartResponse{
		Part: repoConverter.GetPartResponseToModel(*part),
	}, nil

}

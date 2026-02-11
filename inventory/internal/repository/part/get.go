package part

import (
	"context"

	repoConverter "github.com/PhilSuslov/homework/inventory/internal/repository/converter"
	"github.com/PhilSuslov/homework/inventory/internal/model"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

func (r *repository) GetPart(ctx context.Context, info model.GetPartRequest) (model.GetPartResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[info.Uuid]
	if !ok {
		return model.GetPartResponse{}, repoModel.ErrNotFound
	}

	return model.GetPartResponse{
		Part: repoConverter.GetPartResponseToModel(part),
	}, nil

}

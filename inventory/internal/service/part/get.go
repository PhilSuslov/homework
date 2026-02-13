package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

func (s *service) GetPart(ctx context.Context, info model.GetPartRequest) (model.GetPartResponse, error) {
	part, err := s.repository.GetPart(ctx, info)

	if err != nil {
		return model.GetPartResponse{}, repoModel.ErrNotFound
	}

	return part, nil

}

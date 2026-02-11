package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
)

func (s *service) ListParts(ctx context.Context, req model.ListPartsRequest) (model.ListPartsResponse, error) {
	part, err := s.repository.ListParts(ctx, req)

	if err != nil {
		return model.ListPartsResponse{}, err
	}

	return part, nil
}

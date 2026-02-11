package part

import (
	"context"
	"log"

	"github.com/PhilSuslov/homework/inventory/internal/model"
)

func (s *service) ListParts(ctx context.Context, req model.ListPartsRequest) (model.ListPartsResponse, error) {
	part, err := s.repository.ListParts(ctx, req)
	log.Println("List part is Service:", part)
	if err != nil {
		return model.ListPartsResponse{}, err
	}

	return part, nil
}

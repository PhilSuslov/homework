package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	serviceConv "github.com/PhilSuslov/homework/inventory/internal/service/converter"
)

func (s *service) ListParts(ctx context.Context, req model.ListPartsRequest) (model.ListPartsResponse, error) {
	part, err := s.repository.ListParts(ctx, serviceConv.ListPartsRequestToRepoModel(req))
	if err != nil {
		return model.ListPartsResponse{}, err
	}

	return model.ListPartsResponse{
		Parts: serviceConv.ListPartsResponseNotesToNote(part.Parts),
	}, nil
}

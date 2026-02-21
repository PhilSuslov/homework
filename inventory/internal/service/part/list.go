package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	serviceConv "github.com/PhilSuslov/homework/inventory/internal/service/converter"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) ListParts(ctx context.Context, req model.ListPartsRequest) (model.ListPartsResponse, error) {
	part, err := s.repository.ListParts(ctx, serviceConv.ListPartsRequestToRepoModel(req))
	if err != nil {
		logger.Error(ctx, "failed to ListPart", zap.Error(err))
		return model.ListPartsResponse{}, err
	}

	return model.ListPartsResponse{
		Parts: serviceConv.ListPartsResponseNotesToNote(part.Parts),
	}, nil
}

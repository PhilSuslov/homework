package part

import (
	"context"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	serviceConv "github.com/PhilSuslov/homework/inventory/internal/service/converter"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) GetPart(ctx context.Context, info model.GetPartRequest) (model.GetPartResponse, error) {
	part, err := s.repository.GetPart(ctx, serviceConv.GetPartRequestToModel(info))

	if err != nil {
		logger.Error(ctx, "failed to get part", zap.String("info.Uuid", info.Uuid), zap.Error(err))
		return model.GetPartResponse{}, model.ErrNotFound
	}
	logger.Info(ctx, "PASS GET PART SERVICE", zap.String("part.Part.OrderUUID,", part.Part.OrderUUID))
	// log.Printf(part.Part.OrderUUID, "PASS GET PART Service")
	return serviceConv.GetPartResponseNoteToModel(part), nil

}

package part

import (
	"context"
	"log"

	"github.com/PhilSuslov/homework/inventory/internal/model"
	serviceConv "github.com/PhilSuslov/homework/inventory/internal/service/converter"
)

func (s *service) GetPart(ctx context.Context, info model.GetPartRequest) (model.GetPartResponse, error) {
	part, err := s.repository.GetPart(ctx, serviceConv.GetPartRequestToModel(info))

	if err != nil {
		return model.GetPartResponse{}, model.ErrNotFound
	}
	log.Printf(part.Part.OrderUUID, "PASS GET PART Service")
	return serviceConv.GetPartResponseNoteToModel(part), nil

}

package part

import (
	"context"
	"errors"

	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (r *NoteRepository) GetPart(ctx context.Context, info repoModel.GetPartRequest) (repoModel.GetPartResponse, error) {
	var note repoModel.Note
	err := r.collection.FindOne(ctx, bson.M{"body.uuid": info.Uuid}).Decode(&note)
	
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, "failed to GetPart", zap.Error(repoModel.ErrNotFound))
			return repoModel.GetPartResponse{}, repoModel.ErrNotFound
		}
		logger.Error(ctx, "failed to GetPart", zap.Error(err))
		return repoModel.GetPartResponse{}, err
	}

	return repoModel.GetPartResponse{
		Part: note,
	}, err

}

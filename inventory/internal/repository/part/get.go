package part

import (
	"context"
	"errors"
	"log"

	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *NoteRepository) GetPart(ctx context.Context, info repoModel.GetPartRequest) (repoModel.GetPartResponse, error) {
	var note repoModel.Note
	// log.Println(info.Uuid)
	// log.Println(bson.M{"body.uuid": info.Uuid})
	err := r.collection.FindOne(ctx, bson.M{"body.uuid": info.Uuid}).Decode(&note)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repoModel.GetPartResponse{}, repoModel.ErrNotFound
		}
		return repoModel.GetPartResponse{}, err
	}
	log.Println("------- GetPart NoteRepo PASS ------------")
	return repoModel.GetPartResponse{
		Part: note,
	}, err

}

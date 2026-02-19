package part

import (
	"go.mongodb.org/mongo-driver/mongo"

	def "github.com/PhilSuslov/homework/inventory/internal/repository"
)

var _ def.InventoryRepository = (*NoteRepository)(nil)

type NoteRepository struct {
	collection *mongo.Collection
	// Parts map[string]*repoModel.Part
}

func NewNoteRepository(collection *mongo.Collection) *NoteRepository {
	// 	collection := db.Collection("notes")
	//
	// 	indexModels := []mongo.IndexModel{
	// 		{
	// 			Keys:    bson.D{{Key: "body.name", Value: 1}},
	// 			Options: options.Index().SetUnique(false),
	// 		},
	// 	}
	//
	// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()
	//
	// 	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	repo := &NoteRepository{
		collection: collection,
	}


	return repo
}

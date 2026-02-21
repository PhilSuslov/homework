package part

import (
	"go.mongodb.org/mongo-driver/mongo"

	def "github.com/PhilSuslov/homework/inventory/internal/repository"
)

var _ def.InventoryRepository = (*NoteRepository)(nil)

type NoteRepository struct {
	collection *mongo.Collection
}

func NewNoteRepository(collection *mongo.Collection) *NoteRepository {

	repo := &NoteRepository{
		collection: collection,
	}

	return repo
}

package part

import (
	"sync"

	"github.com/google/uuid"

	def "github.com/PhilSuslov/homework/inventory/internal/repository"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

var _ def.InventoryRepository = (*Repository)(nil)

type Repository struct {
	mu    sync.RWMutex
	Parts map[string]*repoModel.Part
}

func NewRepository() *Repository {
	repo := &Repository{
		Parts: make(map[string]*repoModel.Part),
	}

	id := uuid.New().String()
	repo.Parts[id] = &repoModel.Part{
		Uuid:  id,
		Name:  "Test Part",
		Price: 100.0,
	}
	return repo
}

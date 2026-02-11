package part

import (
	"sync"

	"github.com/google/uuid"

	def "github.com/PhilSuslov/homework/inventory/internal/repository"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"

)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct{
	mu sync.RWMutex
	parts map[string] *repoModel.Part
}

func NewRepository() *repository {
	repo := &repository{
		parts: make(map[string]*repoModel.Part),
	}

	id := uuid.New().String()
    repo.parts[id] = &repoModel.Part{
        Uuid:  id,
        Name:  "Test Part",
        Price: 100.0,
	}
	return repo
}
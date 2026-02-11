package part

import (
	"sync"

	def "github.com/PhilSuslov/homework/inventory/internal/repository"
	repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"

)

var _ def.InventoryRepository = (*repository)(nil)

type repository struct{
	mu sync.RWMutex
	parts map[string]repoModel.Part
}

func NewRepository() *repository {
	return &repository{
		parts: make(map[string]repoModel.Part),
	}
}
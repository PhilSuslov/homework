package part

import (
	def "github.com/PhilSuslov/homework/inventory/internal/service"
	repo "github.com/PhilSuslov/homework/inventory/internal/repository"
	// repoModel "github.com/PhilSuslov/homework/inventory/internal/repository/model"
)

var _ def.InventoryService = (*service)(nil)

type service struct{
	repository repo.InventoryRepository
}

func NewService(repository repo.InventoryRepository) *service {
	return &service{
		repository: repository,
	}
}
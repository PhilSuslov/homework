package part

import (
	repo "github.com/PhilSuslov/homework/inventory/internal/repository"
	def "github.com/PhilSuslov/homework/inventory/internal/service"
)

var _ def.InventoryService = (*service)(nil)

type service struct {
	repository repo.InventoryRepository
}

func NewService(repository repo.InventoryRepository) *service {
	return &service{
		repository: repository,
	}
}

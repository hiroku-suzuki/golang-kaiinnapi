package repositories

import "kaiin-mock/models"

type FSRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.FSModel, bool)
}

type InMemoryFSRepository struct{}

func NewInMemoryFSRepository() FSRepository {
	return &InMemoryFSRepository{}
}

func (r *InMemoryFSRepository) FindByKaiinNo(kaiinNo string) (*models.FSModel, bool) {
	m, ok := FSByKaiinNo[kaiinNo]
	return m, ok
}

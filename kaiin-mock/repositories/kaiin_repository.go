package repositories

import "kaiin-mock/models"

type KaiinRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.KaiinModel, error)
}

type InMemoryKaiinRepository struct {
	data map[string]*models.KaiinModel
}

func NewInMemoryKaiinRepository() *InMemoryKaiinRepository {
	return &InMemoryKaiinRepository{
		data: mockKaiinData,
	}
}

func (r *InMemoryKaiinRepository) FindByKaiinNo(kaiinNo string) (*models.KaiinModel, error) {
	v, ok := r.data[kaiinNo]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

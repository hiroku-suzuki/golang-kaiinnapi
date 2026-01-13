package repositories

import "kaiin-mock/models"

type KaiintenShokaiKeiyakuRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.KaiintenShokaiKeiyakuByKaiinNoModel, bool)
}

type InMemoryKaiintenShokaiKeiyakuRepository struct{}

func NewInMemoryKaiintenShokaiKeiyakuRepository() KaiintenShokaiKeiyakuRepository {
	return &InMemoryKaiintenShokaiKeiyakuRepository{}
}

func (r *InMemoryKaiintenShokaiKeiyakuRepository) FindByKaiinNo(kaiinNo string) (*models.KaiintenShokaiKeiyakuByKaiinNoModel, bool) {
	m, ok := kaiintenShokaiKeiyakuByKaiinNo[kaiinNo]
	return m, ok
}

package repositories

import (
	"fmt"
	"kaiin-mock/models"
)

type KokyakuKanriKeiyakuRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.KokyakuKanriKeiyakuModel, error)
}

type InMemoryKokyakuKanriKeiyakuRepository struct{}

func NewInMemoryKokyakuKanriKeiyakuRepository() *InMemoryKokyakuKanriKeiyakuRepository {
	return &InMemoryKokyakuKanriKeiyakuRepository{}
}

func (r *InMemoryKokyakuKanriKeiyakuRepository) FindByKaiinNo(kaiinNo string) (*models.KokyakuKanriKeiyakuModel, error) {
	m, ok := kokyakuKanriKeiyakuByKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return m, nil
}

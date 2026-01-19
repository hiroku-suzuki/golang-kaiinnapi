package repositories

import (
	"fmt"
	"kaiin-mock/models"
)

type BukkakuRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.BukkakuModel, error)
}

type InMemoryBukkakuRepository struct{}

func NewInMemoryBukkakuRepository() *InMemoryBukkakuRepository {
	return &InMemoryBukkakuRepository{}
}

func (r *InMemoryBukkakuRepository) FindByKaiinNo(kaiinNo string) (*models.BukkakuModel, error) {
	m, ok := bukkakuByKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return m, nil
}
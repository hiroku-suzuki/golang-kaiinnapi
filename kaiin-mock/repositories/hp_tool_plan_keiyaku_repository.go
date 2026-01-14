package repositories

import (
	"fmt"
	"kaiin-mock/models"
)

type HpToolPlanKeiyakuRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.HpToolPlanKeiyakuModel, error)
}

type InMemoryHpToolPlanKeiyakuRepository struct{}

func NewInMemoryHpToolPlanKeiyakuRepository() *InMemoryHpToolPlanKeiyakuRepository {
	return &InMemoryHpToolPlanKeiyakuRepository{}
}

func (r *InMemoryHpToolPlanKeiyakuRepository) FindByKaiinNo(kaiinNo string) (*models.HpToolPlanKeiyakuModel, error) {
	m, ok := hpToolPlanKeiyakuByKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return m, nil
}

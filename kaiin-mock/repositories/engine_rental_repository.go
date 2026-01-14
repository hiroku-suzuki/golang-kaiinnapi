package repositories

import (
	"fmt"
	"kaiin-mock/models"
)

type EngineRentalRepository interface {
	FindByControlPanelKaiinNo(kaiinNo string) (*models.EngineRentalModel, error)
}

type InMemoryEngineRentalRepository struct{}

func NewInMemoryEngineRentalRepository() *InMemoryEngineRentalRepository {
	return &InMemoryEngineRentalRepository{}
}

func (r *InMemoryEngineRentalRepository) FindByControlPanelKaiinNo(kaiinNo string) (*models.EngineRentalModel, error) {
	m, ok := engineRentalByControlPanelKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return m, nil
}

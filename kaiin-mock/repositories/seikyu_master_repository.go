package repositories

import (
	"fmt"
	"kaiin-mock/models"
)

type SeikyuMasterRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.SeikyuMasterModel, error)
}

type InMemorySeikyuMasterRepository struct{}

func NewInMemorySeikyuMasterRepository() *InMemorySeikyuMasterRepository {
	return &InMemorySeikyuMasterRepository{}
}

func (r *InMemorySeikyuMasterRepository) FindByKaiinNo(kaiinNo string) (*models.SeikyuMasterModel, error) {
	m, ok := seikyuMasterByKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return m, nil
}
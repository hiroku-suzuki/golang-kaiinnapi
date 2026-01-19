package repositories

import (
	"fmt"
	"kaiin-mock/models"
)

type CreditSettlementRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.CreditSettlementModel, error)
}

type InMemoryCreditSettlementRepository struct{}

func NewInMemoryCreditSettlementRepository() *InMemoryCreditSettlementRepository {
	return &InMemoryCreditSettlementRepository{}
}

func (r *InMemoryCreditSettlementRepository) FindByKaiinNo(kaiinNo string) (*models.CreditSettlementModel, error) {
	m, ok := creditSettlementByKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return m, nil
}
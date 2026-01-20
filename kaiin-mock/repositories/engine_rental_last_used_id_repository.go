package repositories

import "fmt"

type EngineRentalLastUsedEngineRentalIdRepository interface {
	FindByKaiinNo(kaiinNo string) (string, error)
}

type InMemoryEngineRentalLastUsedEngineRentalIdRepository struct{}

func NewInMemoryEngineRentalLastUsedEngineRentalIdRepository() *InMemoryEngineRentalLastUsedEngineRentalIdRepository {
	return &InMemoryEngineRentalLastUsedEngineRentalIdRepository{}
}

func (r *InMemoryEngineRentalLastUsedEngineRentalIdRepository) FindByKaiinNo(kaiinNo string) (string, error) {
	v, ok := lastUsedEngineRentalIdByKaiinNo[kaiinNo]
	if !ok {
		return "", fmt.Errorf("not found")
	}
	return v, nil
}

package services

import "kaiin-mock/repositories"

type EngineRentalLastUsedEngineRentalIdService struct {
	repo repositories.EngineRentalLastUsedEngineRentalIdRepository
}

func NewEngineRentalLastUsedEngineRentalIdService(repo repositories.EngineRentalLastUsedEngineRentalIdRepository) *EngineRentalLastUsedEngineRentalIdService {
	return &EngineRentalLastUsedEngineRentalIdService{repo: repo}
}

func (s *EngineRentalLastUsedEngineRentalIdService) GetByKaiinNo(kaiinNo string) (string, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

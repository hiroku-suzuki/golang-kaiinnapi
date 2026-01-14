package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type EngineRentalService struct {
	repo repositories.EngineRentalRepository
}

func NewEngineRentalService(repo repositories.EngineRentalRepository) *EngineRentalService {
	return &EngineRentalService{repo: repo}
}

// queryパラメータは今後の拡張のために受け取る（現時点では空モック返却なので未使用）
func (s *EngineRentalService) GetByControlPanelKaiinNo(kaiinNo string, includeHistory bool, includeReserve bool, validDate string) (*models.EngineRentalModel, error) {
	return s.repo.FindByControlPanelKaiinNo(kaiinNo)
}

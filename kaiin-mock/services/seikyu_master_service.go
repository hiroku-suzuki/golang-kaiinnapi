package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type SeikyuMasterService struct {
	repo repositories.SeikyuMasterRepository
}

func NewSeikyuMasterService(repo repositories.SeikyuMasterRepository) *SeikyuMasterService {
	return &SeikyuMasterService{repo: repo}
}

func (s *SeikyuMasterService) GetByKaiinNo(kaiinNo string) (*models.SeikyuMasterModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}
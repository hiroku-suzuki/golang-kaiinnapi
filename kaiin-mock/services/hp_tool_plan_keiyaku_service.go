package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type HpToolPlanKeiyakuService struct {
	repo repositories.HpToolPlanKeiyakuRepository
}

func NewHpToolPlanKeiyakuService(repo repositories.HpToolPlanKeiyakuRepository) *HpToolPlanKeiyakuService {
	return &HpToolPlanKeiyakuService{repo: repo}
}

func (s *HpToolPlanKeiyakuService) GetByKaiinNo(kaiinNo string) (*models.HpToolPlanKeiyakuModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

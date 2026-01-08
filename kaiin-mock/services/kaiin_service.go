package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type KaiinService interface {
	GetKaiin(kaiinNo string) (*models.KaiinModel, error)
}

type kaiinService struct {
	repo repositories.KaiinRepository
}

func NewKaiinService(repo repositories.KaiinRepository) KaiinService {
	return &kaiinService{repo: repo}
}

func (s *kaiinService) GetKaiin(kaiinNo string) (*models.KaiinModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type KaiintenShokaiKeiyakuService struct {
	repo repositories.KaiintenShokaiKeiyakuRepository
}

func NewKaiintenShokaiKeiyakuService(repo repositories.KaiintenShokaiKeiyakuRepository) *KaiintenShokaiKeiyakuService {
	return &KaiintenShokaiKeiyakuService{repo: repo}
}

func (s *KaiintenShokaiKeiyakuService) FindByKaiinNo(kaiinNo string) (*models.KaiintenShokaiKeiyakuByKaiinNoModel, bool) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

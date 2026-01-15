package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type KokyakuKanriKeiyakuService struct {
	repo repositories.KokyakuKanriKeiyakuRepository
}

func NewKokyakuKanriKeiyakuService(repo repositories.KokyakuKanriKeiyakuRepository) *KokyakuKanriKeiyakuService {
	return &KokyakuKanriKeiyakuService{repo: repo}
}

func (s *KokyakuKanriKeiyakuService) GetByKaiinNo(kaiinNo string) (*models.KokyakuKanriKeiyakuModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

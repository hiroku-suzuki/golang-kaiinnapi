package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type BukkakuService struct {
	repo repositories.BukkakuRepository
}

func NewBukkakuService(repo repositories.BukkakuRepository) *BukkakuService {
	return &BukkakuService{repo: repo}
}

func (s *BukkakuService) GetByKaiinNo(kaiinNo string) (*models.BukkakuModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}
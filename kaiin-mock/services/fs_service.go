package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type FSService struct {
	repo repositories.FSRepository
}

func NewFSService(repo repositories.FSRepository) *FSService {
	return &FSService{repo: repo}
}

func (s *FSService) FindByKaiinNo(kaiinNo string) (*models.FSModel, bool) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

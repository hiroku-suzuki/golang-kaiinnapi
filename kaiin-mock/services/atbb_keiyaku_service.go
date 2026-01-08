package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type AtbbKeiyakuService interface {
	GetAtbbKeiyaku(kaiinNo string) (*models.AtbbKeiyakuModel, error)
}

type atbbKeiyakuService struct {
	repo repositories.AtbbKeiyakuRepository
}

func NewAtbbKeiyakuService(repo repositories.AtbbKeiyakuRepository) AtbbKeiyakuService {
	return &atbbKeiyakuService{repo: repo}
}

func (s *atbbKeiyakuService) GetAtbbKeiyaku(kaiinNo string) (*models.AtbbKeiyakuModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

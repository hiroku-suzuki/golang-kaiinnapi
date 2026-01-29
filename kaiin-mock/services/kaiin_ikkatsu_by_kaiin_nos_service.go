package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type KaiinIkkatsuByKaiinNosService struct {
	repo repositories.KaiinIkkatsuByKaiinNosRepository
}

func NewKaiinIkkatsuByKaiinNosService(repo repositories.KaiinIkkatsuByKaiinNosRepository) *KaiinIkkatsuByKaiinNosService {
	return &KaiinIkkatsuByKaiinNosService{repo: repo}
}

func (s *KaiinIkkatsuByKaiinNosService) GetByKaiinNos(kaiinNos string) ([]models.KaiinIkkatsuItem, error) {
	return s.repo.FindByKaiinNos(kaiinNos)
}

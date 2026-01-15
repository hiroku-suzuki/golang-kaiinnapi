package services

import "kaiin-mock/repositories"

type KaiinListSameTenpoKaiinNoService struct {
	repo repositories.KaiinListSameTenpoKaiinNoRepository
}

func NewKaiinListSameTenpoKaiinNoService(repo repositories.KaiinListSameTenpoKaiinNoRepository) *KaiinListSameTenpoKaiinNoService {
	return &KaiinListSameTenpoKaiinNoService{repo: repo}
}

func (s *KaiinListSameTenpoKaiinNoService) GetByKaiinNo(kaiinNo string) ([]string, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}

package services

import (
	"kaiin-mock/models"
	"kaiin-mock/repositories"
)

type CreditSettlementService struct {
	repo repositories.CreditSettlementRepository
}

func NewCreditSettlementService(repo repositories.CreditSettlementRepository) *CreditSettlementService {
	return &CreditSettlementService{repo: repo}
}

func (s *CreditSettlementService) GetByKaiinNo(kaiinNo string) (*models.CreditSettlementModel, error) {
	return s.repo.FindByKaiinNo(kaiinNo)
}
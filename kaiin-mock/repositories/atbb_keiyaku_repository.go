package repositories

import (
	"errors"

	"kaiin-mock/models"
)

var ErrNotFound = errors.New("not found")

type AtbbKeiyakuRepository interface {
	FindByKaiinNo(kaiinNo string) (*models.AtbbKeiyakuModel, error)
}

// In-memory 実装

type InMemoryAtbbKeiyakuRepository struct {
	data map[string]*models.AtbbKeiyakuModel
}

func NewInMemoryAtbbKeiyakuRepository() *InMemoryAtbbKeiyakuRepository {
	return &InMemoryAtbbKeiyakuRepository{
		// モックデータはパッケージ変数（mockdata.go）で定義
		data: mockAtbbKeiyakuData,
	}
}

func (r *InMemoryAtbbKeiyakuRepository) FindByKaiinNo(kaiinNo string) (*models.AtbbKeiyakuModel, error) {
	v, ok := r.data[kaiinNo]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

package repositories

import "kaiin-mock/models"

type KaiinIkkatsuByKaiinNosRepository interface {
	FindByKaiinNos(kaiinNos string) ([]models.KaiinIkkatsuItem, error)
}

type InMemoryKaiinIkkatsuByKaiinNosRepository struct{}

func NewInMemoryKaiinIkkatsuByKaiinNosRepository() *InMemoryKaiinIkkatsuByKaiinNosRepository {
	return &InMemoryKaiinIkkatsuByKaiinNosRepository{}
}

func (r *InMemoryKaiinIkkatsuByKaiinNosRepository) FindByKaiinNos(kaiinNos string) ([]models.KaiinIkkatsuItem, error) {
	// 「空でOK」なので、キーが無くても空スライスを返す（404にしない）
	items, ok := kaiinIkkatsuByKaiinNos[kaiinNos]
	if !ok {
		return []models.KaiinIkkatsuItem{}, nil
	}
	// nil だと JSON が null になり得るので空スライスに寄せる
	if items == nil {
		return []models.KaiinIkkatsuItem{}, nil
	}
	return items, nil
}

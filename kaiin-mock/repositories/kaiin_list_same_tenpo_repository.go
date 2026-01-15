package repositories

import "fmt"

// 同一店舗会員番号のリスト（文字列配列）を返す

type KaiinListSameTenpoKaiinNoRepository interface {
	FindByKaiinNo(kaiinNo string) ([]string, error)
}

type InMemoryKaiinListSameTenpoKaiinNoRepository struct{}

func NewInMemoryKaiinListSameTenpoKaiinNoRepository() *InMemoryKaiinListSameTenpoKaiinNoRepository {
	return &InMemoryKaiinListSameTenpoKaiinNoRepository{}
}

func (r *InMemoryKaiinListSameTenpoKaiinNoRepository) FindByKaiinNo(kaiinNo string) ([]string, error) {
	v, ok := sameTenpoKaiinNoByKaiinNo[kaiinNo]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return v, nil
}

package repository

import (
	"sort"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
)

type InMemoryTransacRepository struct {
	store  map[int]*domain.Transaction
	nextID int
}

func NewInMemoryTransacRepository() *InMemoryTransacRepository {
	return &InMemoryTransacRepository{store: make(map[int]*domain.Transaction)}
}

func (r *InMemoryTransacRepository) Create(t *domain.Transaction) error {
	r.nextID++
	t.ID = r.nextID
	r.store[t.ID] = t
	return nil
}

func (r *InMemoryTransacRepository) FindByAccountId(accountId int) []*domain.Transaction {
	var result []*domain.Transaction

	for _, t := range r.store {
		if t.AccountID == accountId {
			result = append(result, t)
		}
	}

	// sort
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID // use '>' for descending
	})

	return result
}

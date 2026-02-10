package repository

import "github.com/rodrigo-militao/pismo-tech-case/internal/domain"

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

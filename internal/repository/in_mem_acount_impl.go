package repository

import "github.com/rodrigo-militao/pismo-tech-case/internal/domain"

type InMemoryAccountRepository struct {
	store  map[int]*domain.Account
	nextID int
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{store: make(map[int]*domain.Account)}
}

func (r *InMemoryAccountRepository) Create(a *domain.Account) error {
	r.nextID++
	a.ID = r.nextID
	r.store[a.ID] = a
	return nil
}

func (r *InMemoryAccountRepository) FindByID(id int) (*domain.Account, error) {
	return r.store[id], nil
}

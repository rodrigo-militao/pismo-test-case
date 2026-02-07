package repository

import "github.com/rodrigo-militao/pismo-tech-case/internal/domain"

type AccountRepository interface {
	Create(account *domain.Account) error
	FindByID(id int) (*domain.Account, error)
}

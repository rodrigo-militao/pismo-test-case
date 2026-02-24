package repository

import "github.com/rodrigo-militao/pismo-tech-case/internal/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
	FindByAccountId(accountId int) []*domain.Transaction
}

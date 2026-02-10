package usecase

import (
	"errors"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
)

type CreateTransactionInput struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	transRepo   repository.TransactionRepository
	accountRepo repository.AccountRepository
}

func NewCreateTransactionUseCase(tr repository.TransactionRepository, ar repository.AccountRepository) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transRepo:   tr,
		accountRepo: ar,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInput) (*domain.Transaction, error) {
	account, err := uc.accountRepo.FindByID(input.AccountID)
	if account == nil || err != nil {
		return nil, errors.New("account not found")
	}

	transaction, err := domain.NewTransaction(
		input.AccountID,
		input.OperationTypeID,
		input.Amount,
	)

	if err != nil {
		return nil, err
	}

	err = uc.transRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

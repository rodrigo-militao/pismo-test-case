package usecase

import (
	"errors"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
)

type CreateAccountUseCase struct {
	repo repository.AccountRepository
}

func NewCreateAccountUseCase(repo repository.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{repo: repo}
}

func (uc *CreateAccountUseCase) Execute(documentNumber string) (*domain.Account, error) {
	if documentNumber == "" {
		return nil, errors.New("Document number cannot be empty")
	}

	account := &domain.Account{
		DocumentNumber: documentNumber,
	}

	err := uc.repo.Create(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

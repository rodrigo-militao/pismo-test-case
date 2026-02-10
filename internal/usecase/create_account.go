package usecase

import (
	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
)

type CreateAccountInput struct {
	DocumentNumber string `json:"document_number"`
}
type CreateAccountUseCase struct {
	repo repository.AccountRepository
}

func NewCreateAccountUseCase(repo repository.AccountRepository) *CreateAccountUseCase {
	return &CreateAccountUseCase{repo: repo}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInput) (*domain.Account, error) {
	account, err := domain.NewAccount(input.DocumentNumber)
	if err != nil {
		return nil, err
	}

	err = uc.repo.Create(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

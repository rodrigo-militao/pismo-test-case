package usecase

import (
	"errors"
	"strconv"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
)

type GetAccountUseCase struct {
	repo repository.AccountRepository
}

func NewGetAccountUseCase(repo repository.AccountRepository) *GetAccountUseCase {
	return &GetAccountUseCase{repo: repo}
}

func (uc *GetAccountUseCase) Execute(accountId string) (*domain.Account, error) {
	if accountId == "" {
		return nil, errors.New("accountId cannot be empty")
	}

	intAccountId, err := strconv.Atoi(accountId)
	if err != nil {
		return nil, errors.New("accountId must be int")
	}

	account, err := uc.repo.FindByID(intAccountId)
	if err != nil {
		return nil, errors.New("Unexpected error finding account")
	}

	return account, nil
}

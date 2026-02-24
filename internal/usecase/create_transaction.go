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

	if transaction.IsCredit {
		// lógica de discharge
		remainingPayment := transaction.Amount

		pastTransactions := uc.transRepo.FindByAccountId(account.ID)

		for _, pastTransaction := range pastTransactions {

			if pastTransaction.Balance < 0 && remainingPayment > 0 {

				// converte pra positivo
				debt := -pastTransaction.Balance

				if remainingPayment >= debt {
					// pagamento atual cobre o debito inteiro
					remainingPayment -= debt
					pastTransaction.Balance = 0
				} else {
					// pagamento cobre apenas uma parte
					// ex: -23.5 + 10 = -13.5
					pastTransaction.Balance += remainingPayment
					remainingPayment = 0
				}
			}

			if remainingPayment == 0 {
				break
			}
		}

		transaction.Balance = remainingPayment
	} else {
		// se for compra
		transaction.Balance = transaction.Amount
	}

	if err != nil {
		return nil, err
	}

	err = uc.transRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

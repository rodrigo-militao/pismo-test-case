package usecase_test

import (
	"errors"
	"testing"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(t *domain.Transaction) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindByAccountId(accountId int) []*domain.Transaction {
	// args := m.Called(accountId)
	return []*domain.Transaction{}
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name           string
		input          usecase.CreateTransactionInput
		setupMocks     func(ma *MockAccountRepository, mt *MockTransactionRepository)
		expectedError  string
		expectedAmount float64
	}{
		{
			name: "Success - Should create Purchase (Debit) with negative amount",
			input: usecase.CreateTransactionInput{
				AccountID:       1,
				OperationTypeID: int(domain.OperationNormalPurchase),
				Amount:          50.0,
			},
			expectedAmount: -50.0, // Esperado: Negativo
			setupMocks: func(ma *MockAccountRepository, mt *MockTransactionRepository) {
				ma.On("FindByID", 1).Return(&domain.Account{ID: 1}, nil)

				// 2. TransactionRepo deve ser chamado com uma transação de valor -50.0
				// Usamos mock.MatchedBy para validar a regra de negócio dentro do Mock
				mt.On("Create", mock.MatchedBy(func(tr *domain.Transaction) bool {
					return tr.Amount == -50.0 && tr.OperationTypeID == domain.OperationNormalPurchase
				})).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "Success - Should create Payment (Credit) with positive amount",
			input: usecase.CreateTransactionInput{
				AccountID:       1,
				OperationTypeID: int(domain.OperationCreditVoucher),
				Amount:          100.0,
			},
			expectedAmount: 100.0,
			setupMocks: func(ma *MockAccountRepository, mt *MockTransactionRepository) {
				ma.On("FindByID", 1).Return(&domain.Account{ID: 1}, nil)

				mt.On("Create", mock.MatchedBy(func(tr *domain.Transaction) bool {
					return tr.Amount == 100.0 && tr.OperationTypeID == domain.OperationCreditVoucher
				})).Return(nil)
			},
			expectedError: "",
		},
		{
			name: "Error - Should fail if Account does not exist",
			input: usecase.CreateTransactionInput{
				AccountID:       999,
				OperationTypeID: 1,
				Amount:          50.0,
			},
			setupMocks: func(ma *MockAccountRepository, mt *MockTransactionRepository) {
				ma.On("FindByID", 999).Return(nil, errors.New("account not found"))
			},
			expectedError: "account not found",
		},
		{
			name: "Error - Should fail if Operation Type is invalid",
			input: usecase.CreateTransactionInput{
				AccountID:       1,
				OperationTypeID: 99,
				Amount:          50.0,
			},
			setupMocks: func(ma *MockAccountRepository, mt *MockTransactionRepository) {
				ma.On("FindByID", 1).Return(&domain.Account{ID: 1}, nil)
			},
			expectedError: "invalid operation type",
		},
		{
			name: "Error - Should fail if Repository fails to save",
			input: usecase.CreateTransactionInput{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          50.0,
			},
			setupMocks: func(ma *MockAccountRepository, mt *MockTransactionRepository) {
				ma.On("FindByID", 1).Return(&domain.Account{ID: 1}, nil)
				mt.On("Create", mock.Anything).Return(errors.New("database error"))
			},
			expectedError: "database error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockAccountRepo := new(MockAccountRepository)
			mockTransRepo := new(MockTransactionRepository)

			tc.setupMocks(mockAccountRepo, mockTransRepo)

			uc := usecase.NewCreateTransactionUseCase(mockTransRepo, mockAccountRepo)

			// Act
			result, err := uc.Execute(tc.input)

			// Assert
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedAmount, result.Amount)
			}

			mockAccountRepo.AssertExpectations(t)
			mockTransRepo.AssertExpectations(t)
		})
	}
}

func TestCreateTransactionUseCase_PaymentDischarge(t *testing.T) {
	accountRepo := repository.NewInMemoryAccountRepository()
	transRepo := repository.NewInMemoryTransacRepository()

	account := &domain.Account{
		ID:             1,
		DocumentNumber: "123456789",
	}

	accountRepo.Create(account)

	uc := usecase.NewCreateTransactionUseCase(transRepo, accountRepo)

	t1, _ := uc.Execute(usecase.CreateTransactionInput{AccountID: 1, OperationTypeID: 1, Amount: 50})
	assert.Equal(t, -50.0, t1.Balance)

	t2, _ := uc.Execute(usecase.CreateTransactionInput{AccountID: 1, OperationTypeID: 1, Amount: 23.50})
	assert.Equal(t, -23.50, t2.Balance)

	t3, _ := uc.Execute(usecase.CreateTransactionInput{AccountID: 1, OperationTypeID: 1, Amount: 18.7})
	assert.Equal(t, -18.7, t3.Balance)

	t4, _ := uc.Execute(usecase.CreateTransactionInput{AccountID: 1, OperationTypeID: 4, Amount: 60})
	assert.Equal(t, 0.0, t4.Balance)

	history := transRepo.FindByAccountId(account.ID)
	assert.Equal(t, 0.0, history[0].Balance)
	assert.Equal(t, -13.5, history[1].Balance)
	assert.Equal(t, -18.7, history[2].Balance)

	t5, _ := uc.Execute(usecase.CreateTransactionInput{AccountID: 1, OperationTypeID: 4, Amount: 100})
	assert.Equal(t, 67.8, t5.Balance)

}

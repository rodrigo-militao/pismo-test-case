package usecase

import (
	"errors"
	"testing"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) Create(a *domain.Account) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *MockAccountRepository) FindByID(id int) (*domain.Account, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Account), args.Error(1)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name          string
		inputDoc      string
		setupMock     func(m *MockAccountRepository)
		expectedError string
	}{
		{
			name:     "Success - Should create account",
			inputDoc: "12345678900",
			setupMock: func(m *MockAccountRepository) {
				m.On("Create", mock.AnythingOfType("*domain.Account")).Return(nil)
			},
			expectedError: "",
		}, {
			name:     "Error - Should validade empty document",
			inputDoc: "",
			setupMock: func(m *MockAccountRepository) {
				// Should not calll repository
			},
			expectedError: "Document number cannot be empty",
		}, {
			name:     "Error - Repository fails",
			inputDoc: "12345678900",
			setupMock: func(m *MockAccountRepository) {
				m.On("Create", mock.AnythingOfType("*domain.Account")).Return(errors.New("Unexpected error"))
			},
			expectedError: "Unexpected error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(MockAccountRepository)
			tc.setupMock(mockRepo)
			uc := NewCreateAccountUseCase(mockRepo)

			// Act
			result, err := uc.Execute(tc.inputDoc)

			//Assert
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.inputDoc, result.DocumentNumber)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

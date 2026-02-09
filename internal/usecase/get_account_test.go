package usecase

import (
	"errors"
	"testing"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAccountUseCase_Execute(t *testing.T) {
	testCases := []struct {
		name          string
		inputDoc      string
		expectedID    int
		setupMock     func(m *MockAccountRepository)
		expectedError string
	}{
		{
			name:       "Success - Should get account",
			inputDoc:   "1",
			expectedID: 1,
			setupMock: func(m *MockAccountRepository) {
				m.On("FindByID", mock.AnythingOfType("int")).Return(&domain.Account{ID: 1, DocumentNumber: "12345678900"}, nil)
			},
			expectedError: "",
		},
		{
			name:     "Error - Should validate non int accountId",
			inputDoc: "a",
			setupMock: func(m *MockAccountRepository) {
				// should not call repository
			},
			expectedError: "accountId must be int",
		},
		{
			name:     "Error - Should validate accountId empty",
			inputDoc: "",
			setupMock: func(m *MockAccountRepository) {
				// should not call repository
			},
			expectedError: "accountId cannot be empty",
		},
		{
			name:     "Error - Repository fails",
			inputDoc: "1",
			setupMock: func(m *MockAccountRepository) {
				m.On("FindByID", mock.AnythingOfType("int")).Return(nil, errors.New("Unexpected error"))
			},
			expectedError: "Unexpected error finding account",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(MockAccountRepository)
			tc.setupMock(mockRepo)
			uc := NewGetAccountUseCase(mockRepo)

			// Act
			result, err := uc.Execute(tc.inputDoc)

			// Assert
			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.expectedID, result.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

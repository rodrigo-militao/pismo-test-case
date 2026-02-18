package domain_test

import (
	"testing"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
)

func TestNewAccount(t *testing.T) {
	testCases := []struct {
		name           string
		documentNumber string
		expectedError  string
	}{
		{
			name:           "Success - Should create account",
			documentNumber: "123456789",
			expectedError:  "",
		},
		{
			name:           "Error - Should validate documentNumber is not null",
			documentNumber: "",
			expectedError:  "document number cannot be empty",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange

			// Act
			acc, err := domain.NewAccount(tc.documentNumber)

			// Assert
			if tc.expectedError != "" {
				if err == nil || err.Error() != tc.expectedError {
					t.Fatalf("expected error %q, got %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				if acc == nil {
					t.Fatal("expected account to be created")
				}
			}
		})
	}
}

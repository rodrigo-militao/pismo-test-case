package domain_test

import (
	"testing"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
)

func TestOperationTypeValidators(t *testing.T) {
	testCases := []struct {
		name          string
		op            int
		expectedValid bool
		expectedDebit bool
	}{
		{name: "Normal Purchase", op: int(domain.OperationNormalPurchase), expectedValid: true, expectedDebit: true},
		{name: "Purchase Installments", op: int(domain.OperationPurchaseInstallments), expectedValid: true, expectedDebit: true},
		{name: "Withdrawal", op: int(domain.OperationWithdrawal), expectedValid: true, expectedDebit: true},
		{name: "Credit Voucher", op: int(domain.OperationCreditVoucher), expectedValid: true, expectedDebit: false},
		{name: "Invalid Operation", op: 99, expectedValid: false, expectedDebit: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ot := domain.OperationType(tc.op)

			if got := ot.IsValid(); got != tc.expectedValid {
				t.Fatalf("%s: expected IsValid=%v, got %v", tc.name, tc.expectedValid, got)
			}

			if got := ot.IsDebit(); got != tc.expectedDebit {
				t.Fatalf("%s: expected IsDebit=%v, got %v", tc.name, tc.expectedDebit, got)
			}
		})
	}
}

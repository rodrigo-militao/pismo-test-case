package domain_test

import (
	"testing"
	"time"

	"github.com/rodrigo-militao/pismo-tech-case/internal/domain"
)

func TestNewTransaction_SignAndFields(t *testing.T) {
	testCases := []struct {
		name           string
		accountID      int
		op             int
		rawAmount      float64
		expectedAmount float64
	}{
		{name: "Normal Purchase positive", accountID: 1, op: int(domain.OperationNormalPurchase), rawAmount: 50.0, expectedAmount: -50.0},
		{name: "Normal Purchase negative input", accountID: 1, op: int(domain.OperationNormalPurchase), rawAmount: -50.0, expectedAmount: -50.0},
		{name: "Credit Voucher positive", accountID: 1, op: int(domain.OperationCreditVoucher), rawAmount: 100.0, expectedAmount: 100.0},
		{name: "Credit Voucher negative input", accountID: 1, op: int(domain.OperationCreditVoucher), rawAmount: -100.0, expectedAmount: 100.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tr, err := domain.NewTransaction(tc.accountID, tc.op, tc.rawAmount)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tr == nil {
				t.Fatal("expected transaction to be created")
			}
			if tr.AccountID != tc.accountID {
				t.Fatalf("expected AccountID %d, got %d", tc.accountID, tr.AccountID)
			}
			if tr.OperationTypeID != domain.OperationType(tc.op) {
				t.Fatalf("expected OperationTypeID %v, got %v", domain.OperationType(tc.op), tr.OperationTypeID)
			}
			if tr.Amount != tc.expectedAmount {
				t.Fatalf("expected Amount %v, got %v", tc.expectedAmount, tr.Amount)
			}
			if tr.EventDate.IsZero() {
				t.Fatalf("expected EventDate to be set")
			}
			if time.Since(tr.EventDate) > 2*time.Second {
				t.Fatalf("EventDate not recent")
			}
		})
	}
}

func TestNewTransaction_InvalidOperation(t *testing.T) {
	_, err := domain.NewTransaction(1, 99, 10.0)
	if err == nil || err.Error() != "invalid operation type" {
		t.Fatalf("expected error 'invalid operation type', got %v", err)
	}
}

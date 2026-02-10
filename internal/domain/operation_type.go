package domain

type OperationType int

const (
	OperationNormalPurchase       OperationType = 1
	OperationPurchaseInstallments OperationType = 2
	OperationWithdrawal           OperationType = 3
	OperationCreditVoucher        OperationType = 4
)

type operationConfig struct {
	Description string
	IsDebit     bool
}

var operationRegistry = map[OperationType]operationConfig{
	OperationNormalPurchase:       {Description: "Normal Purchase", IsDebit: true},
	OperationPurchaseInstallments: {Description: "Purchase with installments", IsDebit: true},
	OperationWithdrawal:           {Description: "Withdrawal", IsDebit: true},
	OperationCreditVoucher:        {Description: "Credit Voucher", IsDebit: false},
}

func (ot OperationType) IsValid() bool {
	_, exists := operationRegistry[ot]
	return exists
}

func (ot OperationType) IsDebit() bool {
	config := operationRegistry[ot]
	return config.IsDebit
}

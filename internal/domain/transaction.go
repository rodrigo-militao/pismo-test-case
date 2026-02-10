package domain

import (
	"errors"
	"math"
	"time"
)

type Transaction struct {
	ID              int           `json:"transaction_id"`
	AccountID       int           `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	EventDate       time.Time     `json:"event_date"`
}

func NewTransaction(accountID int, opType int, rawAmount float64) (*Transaction, error) {
	operationType := OperationType(opType)
	if !operationType.IsValid() {
		return nil, errors.New("invalid operation type")
	}

	finalAmount := math.Abs(rawAmount)

	if operationType.IsDebit() {
		finalAmount = -finalAmount
	}

	return &Transaction{
		AccountID:       accountID,
		OperationTypeID: operationType,
		Amount:          finalAmount,
		EventDate:       time.Now(),
	}, nil
}

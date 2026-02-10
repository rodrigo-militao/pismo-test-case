package domain

import "errors"

type Account struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccount(documentNumber string) (*Account, error) {
	if documentNumber == "" {
		return nil, errors.New("document number cannot be empty")
	}

	return &Account{
		DocumentNumber: documentNumber,
	}, nil
}

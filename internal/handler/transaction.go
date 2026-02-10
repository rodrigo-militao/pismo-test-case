package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
)

type TransactionHandler struct {
	createUseCase *usecase.CreateTransactionUseCase
}

func NewTransactionHandler(uc *usecase.CreateTransactionUseCase) *TransactionHandler {
	return &TransactionHandler{createUseCase: uc}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateTransactionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	transaction, err := h.createUseCase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

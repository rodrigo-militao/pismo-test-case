package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
)

type TransactionHandler struct {
	createUseCase *usecase.CreateTransactionUseCase
	transRepo     repository.TransactionRepository
}

func NewTransactionHandler(uc *usecase.CreateTransactionUseCase, repo repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{createUseCase: uc, transRepo: repo}
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

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue("accountId")
	intAccountId, _ := strconv.Atoi(accountId)
	transactions := h.transRepo.FindByAccountId(intAccountId)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transactions)
}

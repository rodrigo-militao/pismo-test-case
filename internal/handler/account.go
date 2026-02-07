package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
)

type AccountHandler struct {
	createUseCase *usecase.CreateAccountUseCase
	getUseCase    *usecase.GetAccountUseCase
}

func NewAccountHandler(createUC *usecase.CreateAccountUseCase, getUC *usecase.GetAccountUseCase) *AccountHandler {
	return &AccountHandler{
		createUseCase: createUC,
		getUseCase:    getUC,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DocumentNumber string `json:"document_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	account, err := h.createUseCase.Execute(input.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue("accountId")

	account, err := h.getUseCase.Execute(accountId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if account == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

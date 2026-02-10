package main

import (
	"log"
	"net/http"

	"github.com/rodrigo-militao/pismo-tech-case/internal/handler"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
)

func main() {
	accountRepo := repository.NewInMemoryAccountRepository()
	transactionRepo := repository.NewInMemoryTransacRepository()

	createAccountUC := usecase.NewCreateAccountUseCase(accountRepo)
	getAccountUC := usecase.NewGetAccountUseCase(accountRepo)

	createTransactionUC := usecase.NewCreateTransactionUseCase(transactionRepo, accountRepo)

	accountHandler := handler.NewAccountHandler(createAccountUC, getAccountUC)
	transactionHandler := handler.NewTransactionHandler(createTransactionUC)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /accounts", accountHandler.CreateAccount)
	mux.HandleFunc("GET /accounts/{accountId}", accountHandler.GetAccount)

	mux.HandleFunc("POST /transactions", transactionHandler.CreateTransaction)

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/rodrigo-militao/pismo-tech-case/internal/handler"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
)

func main() {
	repo := repository.NewInMemoryAccountRepository()

	createAccountUC := usecase.NewCreateAccountUseCase(repo)
	getAccountUC := usecase.NewGetAccountUseCase(repo)

	accountHandler := handler.NewAccountHandler(createAccountUC, getAccountUC)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /accounts", accountHandler.CreateAccount)
	mux.HandleFunc("GET /accounts/{accountId}", accountHandler.GetAccount)

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

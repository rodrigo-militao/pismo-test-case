package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("🚀 Server running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

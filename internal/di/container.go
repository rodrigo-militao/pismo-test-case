package di

import (
	"github.com/rodrigo-militao/pismo-tech-case/internal/handler"
	"github.com/rodrigo-militao/pismo-tech-case/internal/repository"
	"github.com/rodrigo-militao/pismo-tech-case/internal/usecase"
)

type Container struct {
	AccountHandler     *handler.AccountHandler
	TransactionHandler *handler.TransactionHandler
}

func NewContainer() *Container {
	// Repositories
	accountRepo := repository.NewInMemoryAccountRepository()
	transRepo := repository.NewInMemoryTransacRepository()

	// UseCases
	createAccountUC := usecase.NewCreateAccountUseCase(accountRepo)
	getAccountUC := usecase.NewGetAccountUseCase(accountRepo)
	createTransUC := usecase.NewCreateTransactionUseCase(transRepo, accountRepo)

	// Handlers
	accountHandler := handler.NewAccountHandler(createAccountUC, getAccountUC)
	transactionHandler := handler.NewTransactionHandler(createTransUC, transRepo)

	return &Container{
		AccountHandler:     accountHandler,
		TransactionHandler: transactionHandler,
	}
}

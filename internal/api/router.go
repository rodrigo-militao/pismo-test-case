package api

import (
	"net/http"

	"github.com/rodrigo-militao/pismo-tech-case/internal/di"
)

func NewRouter(c *di.Container) *http.ServeMux {
	mux := http.NewServeMux()

	// Accounts
	mux.HandleFunc("POST /accounts", c.AccountHandler.CreateAccount)
	mux.HandleFunc("GET /accounts/{accountId}", c.AccountHandler.GetAccount)

	// Transactions
	mux.HandleFunc("POST /transactions", c.TransactionHandler.CreateTransaction)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}

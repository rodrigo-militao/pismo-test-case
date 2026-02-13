package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rodrigo-militao/pismo-tech-case/internal/api"
	"github.com/rodrigo-militao/pismo-tech-case/internal/di"
)

func main() {
	// Dependency Injection
	container := di.NewContainer()

	// Routing
	router := api.NewRouter(container)

	// Setup Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("🚀 Server running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	api.GracefulShutdown(server)
}

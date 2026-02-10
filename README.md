# Pismo Tech Case - Transaction Service

Este repositório contém a implementação de uma API REST para gestão de contas e transações financeiras. O projeto foi desenvolvido com foco em simplicidade, manutenibilidade e testabilidade, utilizando Go (Golang) e Docker.

## 🛠 Tech Stack & Arquitetura

- **Linguagem:** Go 1.22+
- **Arquitetura:** Clean Architecture (Domain, UseCase, Repository, Handler).
- **Design Patterns:** Factory Method, Repository Pattern, Dependency Injection.
- **Conceitos:** DDD (Rich Domain Model), Object Calisthenics (Fail Fast, No Else).
- **Testes:** Table Driven Tests com Mocks (`testify`).

## 🚀 Como Rodar o Projeto

### Pré-requisitos
- **Docker** (para execução containerizada - Recomendado)
- **Go 1.22+** (apenas para execução local sem Docker)

### Opção 1: Via Docker (Universal & Recomendado)
Esta opção garante o ambiente isolado. Funciona em qualquer terminal (PowerShell, CMD, Bash).
O script abaixo constrói a imagem e inicia o container na porta `8080`.

1. Construa a imagem:
```bash
docker build -t pismo-api .
```
2. Execute o container
```bash
docker run --rm -p 8080:8080 pismo-api
```

### Opção 2: Execução Local (Desenvolvimento)
1. Base as dependências:
```bash
go mod tidy
```
2. Execute o container
```bash
go run cmd/api/main.go
```
#### A API estará disponível em: http://localhost:8080
---

### Como rodar os testes unitários
```bash
go test ./... -v
```
- `./...`: Roda em todas as subpastas (recursivo).

- `-v`: Verbose (mostra o nome de cada teste que rodou).
---

### ⚡ Atalhos (Opcional)

Para usuários de ambientes Unix (Linux/Mac/WSL), o projeto inclui um Makefile e scripts para conveniência:
- Make: Execute `make run` ou `make docker-run`.
- Scripts: Execute `./run.sh` ou `./docker-run.sh`.
---
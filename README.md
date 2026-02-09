# Pismo Tech Case - Transaction Service

Este repositório contém a implementação de uma API REST para gestão de contas e transações financeiras. O projeto foi desenvolvido com foco em simplicidade, manutenibilidade e testabilidade, utilizando Go (Golang) e Docker.

## 🚀 Como Rodar o Projeto

Para atender aos critérios de avaliação de "Easy application execution", o projeto possui scripts utilitários na raiz.

### Pré-requisitos
- **Docker** (para execução containerizada - Recomendado)
- **Go 1.22+** (apenas para execução local sem Docker)

### Opção 1: Via Docker (Ambiente Isolado)
Esta é a forma recomendada para avaliação, garantindo que o ambiente seja idêntico ao de desenvolvimento. O script abaixo constrói a imagem e inicia o container na porta `8080`.

```bash
# Dá permissão de execução (necessário apenas na primeira vez)
chmod +x docker-run.sh

# Roda a aplicação via Docker
./docker-run.sh
```
### Opção 2: Execução Local (Desenvolvimento)

```bash
# Dá permissão de execução
chmod +x run.sh

# Instala dependências e roda a aplicação
./run.sh
```
A API estará disponível em: http://localhost:8080

## Como rodar os testes unitários
```bash
go test ./... -v
```
- `./...`: Roda em todas as subpastas (recursivo).

- `-v`: Verbose (mostra o nome de cada teste que rodou).

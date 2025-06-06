# Voting System 🗳️

Um sistema simples de votação construído em Go, utilizando o framework Echo para criação de APIs RESTful.

## Funcionalidades

- Criar candidatos
- Listar todos os candidatos
- Deletar candidatos
- Criar e registrar votos
- Obter resultados da votação

## Estrutura do Projeto
````
├── controllers/ # Lógica dos endpoints (ex: candidate_controller.go)
├── models/ # Estrutura dos dados (Candidate, Vote, etc.)
├── routes/ # Definição das rotas e handlers
├── main.go # Inicialização da aplicação
└── go.mod # Gerenciamento de dependências
````

## Tecnologias

- [Go](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- Banco de dados (possivelmente em memória ou SQLite)

## Como rodar o projeto

```bash
git clone https://github.com/gustavomelo20/voting-system.git
cd voting-system
go run main.go
```

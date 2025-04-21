# Todo API Service - Golang & MongoDB

Esta é uma aplicação simples desenvolvida para aprender conceitos iniciais da criação de uma API em Golang e começar a utilizar bancos de dados não relacionais como MongoDB.

## Sobre o Projeto

O projeto é uma API REST para gerenciamento de tarefas (Todo List) que permite criar, listar, atualizar e deletar tarefas. A aplicação foi construída usando:

- Go (Golang)
- MongoDB como banco de dados
- Gorilla Mux para roteamento HTTP
- Docker para containerização

## Requisitos

- Go 1.15+
- Docker e Docker Compose
- Git

## Como Executar

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/todo-service-go.git
cd todo-service-go
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```
DB_URI=mongodb://mongouser:mongopwd@localhost:27017
DB_NAME=tododb
DB_COLLECTION_NAME=tasks
```

### 3. Inicie o MongoDB com Docker Compose

Execute o comando abaixo para iniciar o contêiner do MongoDB:

```bash
docker-compose up -d
```

Isso iniciará o MongoDB na porta 27017 com as credenciais definidas no arquivo docker-compose.yml.

### 4. Execute a aplicação

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

## Endpoints da API

### Listar todas as tarefas
```
GET /api/task
```

### Obter uma tarefa específica
```
GET /api/task/{id}
```

### Criar uma nova tarefa
```
POST /api/task
```
Exemplo de corpo da requisição:
```json
{
  "title": "Estudar Golang",
  "description": "Aprender sobre APIs REST com Go",
  "due_date": "2023-12-31",
  "status": false
}
```

### Marcar uma tarefa como concluída
```
PUT /api/task/complete/{id}
```

### Desfazer uma tarefa concluída
```
PUT /api/task/undo/{id}
```

### Deletar uma tarefa
```
DELETE /api/task/{id}
```

### Deletar todas as tarefas
```
DELETE /api/task
```

## Estrutura do Projeto

```
todo-service-go/
├── main.go           # Ponto de entrada da aplicação
├── models/
│   └── task.go       # Modelo de dados da tarefa
├── middleware/
│   └── middleware.go # Handlers das rotas e conexão com MongoDB
├── router/
│   └── router.go     # Configuração das rotas da API
├── docker-compose.yml # Configuração do Docker para MongoDB
└── .env              # Arquivo de variáveis de ambiente
```

## Docker Compose

O arquivo `docker-compose.yml` configura o MongoDB:

```yaml
version: '3'

services:
  mongodb:
    image: mongo:latest
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: mongouser
      MONGO_INITDB_ROOT_PASSWORD: mongopwd
```

Para iniciar o contêiner:
```bash
docker-compose up -d
```

Para parar o contêiner:
```bash
docker-compose down
```


Melhorias futuras:
- Adicionar autenticação de usuários
- Implementar validação de dados de entrada
- Adicionar testes unitários e de integração
- Implementar paginação para listagem de tarefas
- Melhorar o tratamento de erros

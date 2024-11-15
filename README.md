# Management API Go 🚀

Bem-vindo ao projeto **Management API Go**! Este é um exemplo de aplicação que demonstra como integrar várias tecnologias modernas em um único projeto. O objetivo é mostrar como construir uma API robusta utilizando Go, MySQL, Kafka e Docker.

## 📝 Índice

- [Sobre o Projeto](#-sobre-o-projeto)
- [Tecnologias Utilizadas](#-tecnologias-utilizadas)
- [Estrutura do Projeto](#-estrutura-do-projeto)
- [Como Executar Localmente](#-como-executar-localmente)
- [Como Contribuir](#-como-contribuir)
- [Licença](#-licença)

## 📖 Sobre o Projeto

Esta aplicação é uma API de gerenciamento desenvolvida em Go, com persistência de dados em MySQL e comunicação assíncrona utilizando Kafka. O projeto visa ilustrar como diferentes tecnologias podem ser aplicadas em conjunto para criar uma solução completa.

## 🛠 Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação principal do projeto.
- **MySQL**: Banco de dados relacional para armazenamento de dados.
- **Kafka**: Sistema de mensagens distribuído para comunicação assíncrona.
- **Docker & Docker Compose**: Containerização dos serviços e configuração facilitada do ambiente.
- **GitHub Actions**: Integração contínua para testes automatizados.
- **Makefile**: Automação de comandos comuns de desenvolvimento.

## 📁 Estrutura do Projeto

```
management-api-go/
├── .github/
│   └── workflows/
│       └── ci.yml
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   │   └── user.go
│   │   └── repositories/
│   │       └── user_repository.go
│   ├── infra/
│   │   ├── database/
│   │   │   └── db.go
│   │   ├── kafka/
│   │   │   └── producer.go
│   │   └── persistence/
│   │       └── user_repository_db.go
│   ├── interfaces/
│   │   ├── controllers/
│   │   │   └── user_controller.go
│   │   └── routes/
│   │       └── routes.go
│   └── usecases/
│       ├── user_usecase.go
│       └── user_usecase_test.go
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── README.md

```

## 🚀 Como Executar Localmente

### Pré-requisitos

- **Go** instalado (versão 1.23 ou superior)
- **Docker** e **Docker Compose** instalados

### Passos

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/dsperax/management-api-go.git
   cd management-api-go
   ```

2. **Inicie os serviços com Docker Compose:**

    ```
    make setup-db
    ```
Isso irá iniciar o MySQL e o Kafka em contêineres Docker.

3. **Execute a aplicação:**

    ```
    make run
    ```
A API estará rodando em http://localhost:8080.

4. **Teste a API:**    

Utilize ferramentas como Postman ou curl para interagir com a API.

## 🤝 Como Contribuir
1- Faça um fork do projeto.
2- Crie uma branch para sua feature (git checkout -b feature/nova-feature).
3- Commit suas alterações (git commit -m 'Adiciona nova feature').
4- Faça o push para a branch (git push origin feature/nova-feature).
5- Abra um Pull Request.
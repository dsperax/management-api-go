# (EN) Management API Go 🚀

Welcome to the **Management API Go** project! This is a sample application that demonstrates how to integrate various modern technologies into a single project. The goal is to show how to build a robust API using Go, MySQL, Kafka, and Docker.

## 📝 Index

- [About the Project](#-sobre-o-projeto)
- [Technologies Used](#-tecnologias-utilizadas)
- [Project Structure](#-estrutura-do-projeto)
- [How to Run Locally](#-como-executar-localmente)
- [How to Contribute](#-como-contribuir)
- [License](#-licença)

## 📖 About the Project

This application is a management API developed in Go, with data persistence in MySQL and asynchronous communication using Kafka. The project aims to illustrate how different technologies can be applied together to create a complete solution.

## 🛠 Technologies Used

- **Go (Golang)**: Main programming language of the project.
- **MySQL**: Relational database for data storage.
- **Kafka**: Distributed messaging system for asynchronous communication.
- **Docker & Docker Compose**: Containerization of services and simplified environment setup.
- **GitHub Actions**: Continuous integration for automated testing.
- **Makefile**: Automation of common development commands.

## 📁 Project Structure

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

## 🚀 How to Run Locally

### Prerequisites

- **Go** installed (version 1.23 or higher)
- **Docker** e **Docker Compose** installed

### Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dsperax/management-api-go.git
   cd management-api-go
   ```

2. **Start the services with Docker Compose:**

    ```
    make setup-db
    ```
This will start MySQL and Kafka in Docker containers.

3. **Run the application:**

    ```
    make run
    ```
The API will be running at http://localhost:8080.

4. **Swagger:**
   
   ```
   http://localhost:8080/swagger/index.html
   ```

5. **Test the API:**

Use tools like Postman or curl to interact with the API.

6. **Grafana:**

    ```
    http://localhost:3000
    ```

## 🤝 How to Contribute
1. Fork the project.
2. Create a branch for your feature (git checkout -b feature/new-feature).
3. Commit your changes (git commit -m 'Add new feature').
4. Push to the branch (git push origin feature/new-feature).
5. Open a Pull Request.
6. Wait for approval.

<hr>

# (PT-BR) Management API Go 🚀

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

4. **Swagger:**
   
   ```
   http://localhost:8080/swagger/index.html
   ```
   
5. **Teste a API:**    

Utilize ferramentas como Postman ou curl para interagir com a API.

6. **Grafana:**

    ```
    http://localhost:3000
    ```

## 🤝 Como Contribuir
1. Faça um fork do projeto.
2. Crie uma branch para sua feature (git checkout -b feature/nova-feature).
3. Commit suas alterações (git commit -m 'Adiciona nova feature').
4. Faça o push para a branch (git push origin feature/nova-feature).
5. Abra um Pull Request.
6. Espere a aprovação.

```
# TODO

### Passo 1: Configurar o Grafana para Usar o Prometheus como Fonte de Dados
Após iniciar os serviços, você precisará configurar o Grafana para usar o Prometheus como fonte de dados.

1.1. Acessar o Grafana
Abra o Grafana no navegador:

arduino
Copiar código
http://localhost:3000
1.2. Login no Grafana
Usuário: admin
Senha: admin
Você será solicitado a alterar a senha após o primeiro login.

1.3. Adicionar Fonte de Dados
Clique em "Add your first data source" ou vá em Configuration (engrenagem) > Data Sources e clique em "Add data source".

Selecione Prometheus.

Configure a fonte de dados:

Name: Prometheus
URL: http://prometheus:9090
Access: Server (Default)
Clique em "Save & Test".

Você deve ver a mensagem "Data source is working".

### Passo 2: Importar um Dashboard para Visualizar as Métricas
Você pode criar seus próprios dashboards ou importar dashboards existentes.

2.1. Importar um Dashboard Existente
No Grafana, clique no ícone de "Plus" na barra lateral e selecione "Import".

Importar pelo ID do Dashboard

Você pode usar o ID de um dashboard existente da comunidade. Por exemplo, para um dashboard básico do Go, use o ID "1860".
Clique em "Load".

Selecione a fonte de dados

Em Prometheus, selecione a fonte de dados que você configurou.
Clique em "Import".

Agora você deve ver um dashboard com métricas da sua aplicação Go.

### Passo 3: Verificar se as Métricas Estão Sendo Coletadas
Certifique-se de que o Prometheus está coletando as métricas da sua aplicação.

3.1. Acessar o Prometheus
Abra o Prometheus no navegador:

arduino
Copiar código
http://localhost:9090
Você pode executar consultas para verificar as métricas disponíveis.

### Passo 8: Testar o Ambiente
Faça algumas requisições à sua aplicação Go para gerar atividade.

Verifique no Grafana se as métricas estão sendo atualizadas nos dashboards.
```
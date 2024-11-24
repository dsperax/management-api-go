.PHONY: build run test lint staticcheck docker-build docker-run compose-up compose-down setup-db clean swagger open-grafana

build:
	go build -o bin/main ./cmd/api/main.go

run:
	go run ./cmd/api/main.go

test:
	go test ./... -v

lint:
	@echo "Executando Golint..."
	@go install golang.org/x/lint/golint@latest
	@OUTPUT=$$(golint ./...); \
	if [ -n "$$OUTPUT" ]; then \
		echo "$$OUTPUT"; \
		exit 1; \
	else \
		echo "Nenhum problema encontrado pelo Golint."; \
	fi

staticcheck:
	@echo "Executando Staticcheck..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@staticcheck ./...

docker-build:
	docker build -t dsperax/management-api-go .

docker-run:
	docker run -p 8080:8080 dsperax/management-api-go

compose-up:
	docker compose up -d

compose-down:
	docker compose down

swagger:
	@swag init -g main.go -o ./docs -d ./cmd/api,./internal

setup-db: compose-up
	@echo "Aguardando o MySQL iniciar..."
	@until docker compose exec mysql mysqladmin ping -uuser -puserpassword --silent; do\
		echo "Aguardando o MySQL ficar disponível...";\
		sleep 2;\
	done
	@echo "Banco de dados inicializado com sucesso."
# OBS: A senha do root do DB jamais deve ficar exposta dessa maneira, isso é apenas um facilitador para o desenvolvimento.

clean:
	@read -p "Tem certeza de que deseja remover todos os containers e volumes do projeto? [y/N] " confirm && \
	if [ "$$confirm" = "y" ]; then \
		docker compose down -v; \
	else \
		echo "Operação cancelada."; \
	fi

open-grafana:
	@echo "Abrindo Grafana em http://localhost:3000"
	xdg-open http://localhost:3000 || open http://localhost:3000

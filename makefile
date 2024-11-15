.PHONY: build run test docker-build docker-run compose-up compose-down setup-db clean

build:
	go build -o bin/main ./cmd/api/main.go

run:
	go run ./cmd/api/main.go

test:
	go test ./... -v

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

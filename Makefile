APP_NAME=trip-weather-deals
MAIN_PACKAGE=./cmd/api
BINARY_NAME=trip-weather-deals

.PHONY: help run test fmt vet build clean docker-build docker-run

help:
	@echo "Comandos disponíveis:"
	@echo "  make run          - Executa a aplicação localmente"
	@echo "  make test         - Executa os testes"
	@echo "  make fmt          - Formata o código"
	@echo "  make vet          - Executa análise estática"
	@echo "  make build        - Gera o binário"
	@echo "  make clean        - Remove o binário gerado"
	@echo "  make docker-build - Gera a imagem Docker"
	@echo "  make docker-run   - Executa o container Docker"

run:
	go run $(MAIN_PACKAGE)

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

clean:
	rm -f $(BINARY_NAME)

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm -p 8080:8080 $(APP_NAME):latest
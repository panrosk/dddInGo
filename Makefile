# Variables
APP_NAME=coworking-app
BUILD_DIR=bin
SRC_DIR=cmd/main.go

# Comandos básicos
.PHONY: all build run test clean watch lint install-air

# Compilar la aplicación
build:
	@echo "🔨 Compilando la aplicación..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Ejecutar la aplicación sin compilar
run:
	@echo "🚀 Ejecutando la aplicación..."
	go run $(SRC_DIR)

# Compilar y ejecutar la aplicación
run-build: build
	@echo "🚀 Ejecutando la aplicación compilada..."
	./$(BUILD_DIR)/$(APP_NAME)

# Ejecutar los tests
test:
	@echo "🧪 Ejecutando tests..."
	gotestsum --format testname

# Ejecutar tests con cobertura
coverage:
	@echo "📊 Ejecutando tests con cobertura..."
	go test -cover ./...

# Ejecutar lint con golangci-lint
lint:
	@echo "🔍 Ejecutando análisis estático..."
	golangci-lint run ./...

# Limpiar binarios
clean:
	@echo "🧹 Limpiando archivos generados..."
	rm -rf $(BUILD_DIR)

# Instalar Air para desarrollo en caliente
install-air:
	@echo "📦 Instalando Air..."
	go install github.com/cosmtrek/air@latest

# Ejecutar el servidor en modo Watch con Air
watch:
	@echo "🔄 Iniciando modo Watch con Air..."
	air -c .air.toml

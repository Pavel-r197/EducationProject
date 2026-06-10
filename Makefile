BINARY_NAME=myapp
BINARY_MIGRATIONS_NAME=migrations
BUILD_DIR=./cmd

.PHONY: all build run

build:
	@echo "Building binary..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

	@echo "Building binary..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

run: build
	@echo "Запуск миграции $(BINARY_MIGRATIONS_NAME)..."
	$(BUILD_DIR)/$(BINARY_MIGRATIONS_NAME)

	@echo "Запуск приложения $(BINARY_NAME)..."
	$(BUILD_DIR)/$(BINARY_NAME)
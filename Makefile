APP_NAME := kauma
TEST_INPUT_DIR := Aufgaben/input
TEST_OUTPUT_DIR := output
TEST_OUTPUT_FILE := $(TEST_OUTPUT_DIR)/results.json
GO_FILES := $(shell find . -name '*.go')

.PHONY: all
all: test-unit test-json

.PHONY: clean
clean:
	@echo "Cleaning output files..."
	@rm -rf $(TEST_OUTPUT_DIR)

.PHONY: build
build:
	@echo "Building the Go application..."
	@go build -o $(APP_NAME) ./cmd/kauma

.PHONY: test-unit
test-unit:
	@echo "Running unit tests..."
	@go test ./... -v -cover
	@if [ $$? -ne 0 ]; then echo "Unit tests failed!" && exit 1; fi

.PHONY: test-json
test-json: build
	@echo "Running JSON test cases..."
	@mkdir -p $(TEST_OUTPUT_DIR)
	@for file in $(TEST_INPUT_DIR)/*; do \
		echo "Processing $$file..."; \
		./$(APP_NAME) $$file > $(TEST_OUTPUT_FILE); \
		if [ $$? -ne 0 ]; then echo "Testcase execution failed for $$file!" && exit 1; fi \
	done
	@echo "All JSON test cases passed."

.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...
	@if [ $$? -ne 0 ]; then echo "Linting failed!" && exit 1; fi

.PHONY: pipeline
pipeline: lint test-unit test-json
	@echo "All pipeline steps passed successfully."

.PHONY: dev
dev: clean build all
	@echo "Development environment setup complete."

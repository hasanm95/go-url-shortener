BINARY_NAME := go-url-shortening

.PHONY: build run

build:
	@go build -o $(BINARY_NAME) ./cmd

run: build
	@./$(BINARY_NAME)

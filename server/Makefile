# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	@go build -o main.exe main.go

# Run the application
run:
	@go run main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

docker-dev:
	@echo "Building Docker image..."
	@docker compose --profile dev up -d 

# # Live Reload
# watch:
#     if not exist "$(GOPATH)\bin\air.exe" (
#         go install github.com/air-verse/air@latest
#     )
#     "$(GOPATH)\bin\air.exe"

# .PHONY: all build run test clean watch
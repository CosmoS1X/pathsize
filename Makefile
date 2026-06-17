RUN_ARGS ?=
APP_NAME := hexlet-path-size

.PHONY: fmt tidy test lint lint-fix build run vuln clean

.DEFAULT_GOAL := build

fmt:
	golangci-lint fmt

tidy:
	go mod tidy

test: fmt tidy
	go test -v ./...

lint: fmt
	golangci-lint run

lint-fix:
	golangci-lint run --fix

build: lint
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run:
	go run cmd/$(APP_NAME)/main.go $(RUN_ARGS)

vuln:
	govulncheck ./...

clean:
	rm -rf bin/

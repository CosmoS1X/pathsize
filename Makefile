RUN_ARGS ?=
APP_NAME := hexlet-path-size

.PHONY: fmt vet tidy build test run vuln clean

.DEFAULT_GOAL := build

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

tidy:
	go mod tidy

build: vet tidy
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

test: vet tidy
	go test -v ./...

run:
	go run cmd/$(APP_NAME)/main.go $(RUN_ARGS)

vuln:
	govulncheck ./...

clean:
	rm -rf bin/

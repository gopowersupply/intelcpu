help:
	@echo "help           - This help"
	@echo "all            - Run fmt and all quality checkers"
	@echo "fmt            - Make gofmt -s -w ./"
	@echo "test           - Run tests"
	@echo "lint           - Run go lint"
.PHONY: help

all: fmt test lint
.PHONY: all

fmt:
	gofmt -s -w ./
.PHONY: fmt

test:
	go test ./...
.PHONY: test

lint:
	golint ./...
.PHONY: lint
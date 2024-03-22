# Makefile

all: lint test

lint:
	@golangci-lint --version
	golangci-lint run ./...

test:
	go clean -testcache
	go test -v  ./... -race -covermode=atomic -coverprofile=coverage.out -timeout=20m
	@go tool cover -func coverage.out

gen:
	## Generating...
	@go generate ./...

.PHONY: build
build:
	go build  -o ./build/main ./cmd/main.go

run:
	./build/main -json cmd/testdata/t.json

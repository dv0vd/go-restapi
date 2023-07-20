# .DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o ./build -v ./cmd/apiserver 

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
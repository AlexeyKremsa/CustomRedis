.PHONY: test

BINARY_NAME=customredis

build:
	cd ./cmd/restapi && go build -o $(BINARY_NAME) -v

run: build
	cd ./cmd/restapi/ && ./$(BINARY_NAME)

test:
	go test -v ./...
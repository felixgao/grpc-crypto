compile:
	protoc -I=. --go_out=plugins=grpc:./proto *.proto

build:
	go build -o bin/crypto-client ./client

test: examples
	go test ./...

.PHONY: compile build test

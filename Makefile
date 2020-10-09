default: all

.PHONY: build pb

clean:
	rm -rf pkg/pb/books/*.go

pb:
	protoc pkg/pb/books/*.proto --go_out=plugins=grpc:pkg/pb/books

go:
	go build -o ./books-grpc-poc-server ./cmd/server/main.go
	go build -o ./books-grpc-poc-client ./cmd/client/main.go



test: pb
	go test ./...

build: pb go

all: clean test build

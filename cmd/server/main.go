package main

import (
	"flag"
	"github.com/benschw/books-grpc-poc/internal/app"
	"github.com/benschw/books-grpc-poc/internal/fakes"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	addr = flag.String("addr", "localhost:9000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	repo := fakes.NewRepo()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := app.NewServer(repo)

	grpcServer := grpc.NewServer()

	books.RegisterBookServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

package main

import (
	"github.com/benschw/books-grpc-poc/internal/app"
	"github.com/benschw/books-grpc-poc/internal/fakes"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//repo, err := postgres.NewRepo(os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
	//	os.Exit(1)
	//}
	repo := fakes.NewRepo()

	lis, err := net.Listen("tcp", ":9000")
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

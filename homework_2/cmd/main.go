package main

import (
	"log"
	"net"
	"orders-microservice/internal/repository"
	"orders-microservice/internal/service"
	"orders-microservice/pkg/api/test"

	"google.golang.org/grpc"
)

func main() {
	repo := repository.New()
	defer repo.Close()

	ordersService := service.New(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	test.RegisterOrderServiceServer(grpcServer, ordersService)

	grpcServer.Serve(lis)
}

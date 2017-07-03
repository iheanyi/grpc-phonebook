package main

import (
	"log"
	"net"

	"github.com/iheanyi/grpc-phonebook/api"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	svc := api.New()
	api.RegisterPhoneBookServer(srv, svc)
	srv.Serve(lis)
}

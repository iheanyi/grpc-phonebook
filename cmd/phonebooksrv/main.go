package main

import (
	"log"
	"net"

	"github.com/iheanyi/grpc-phonebook/api"
	"github.com/iheanyi/grpc-phonebook/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	svc := server.New()
	api.RegisterPhoneBookServer(srv, svc)
	log.Print("Starting up the server")
	srv.Serve(lis)
}

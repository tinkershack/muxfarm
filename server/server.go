package main

import (
	"log"
	"net"

	mops "github.com/tinkershack/muxfarm/mops"
	pb "github.com/tinkershack/muxfarm/plumber"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50050")
	if err != nil {
		log.Fatalf("Failed to listen on TCP port 50050: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMopsServer(grpcServer, &mops.Server{})
	log.Println("Serving Mops")
	grpcServer.Serve(listen)
}

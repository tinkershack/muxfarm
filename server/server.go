package main

import (
	"log"
	"net"

	"github.com/tinkershack/muxfarm/config"
	"github.com/tinkershack/muxfarm/mops"
	"github.com/tinkershack/muxfarm/plumber"
	"github.com/tinkershack/muxfarm/store"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalf("Fail: acquire config\n%s", err)
	}
	log.Printf("config: %+v", config)
	listen, err := net.Listen("tcp", "localhost:50050")
	if err != nil {
		log.Fatalf("Fail: listen on TCP port 50050: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	ds, err := store.MongoDB(config.MongoDB.URI, config.MongoDB.DBName)
	if err != nil {
		log.Fatalf("Fail: acquire MongoDB object\n%s", err)
	}

	plumber.RegisterMopsServer(grpcServer, mops.Mops(ds))
	log.Println("Serving Mops")
	grpcServer.Serve(listen)
}

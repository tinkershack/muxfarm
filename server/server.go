package main

import (
	"context"
	"log"
	"net"

	"github.com/tinkershack/muxfarm/config"
	"github.com/tinkershack/muxfarm/mops"
	"github.com/tinkershack/muxfarm/plumber"
	"github.com/tinkershack/muxfarm/sense"
	mdb "github.com/tinkershack/muxfarm/store/mongodb"
	"google.golang.org/grpc"
)

// TODO:
// Read values from config
// Moduled structured logging
// Fix error logging

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

	ds, err := mdb.MongoDB(config.MongoDB.URI, config.MongoDB.DBName)
	if err != nil {
		log.Fatalf("Fail: acquire MongoDB object\n%s", err)
	}
	defer ds.Client.Disconnect(context.Background())

	plumber.RegisterMopsServer(grpcServer, mops.Mops(ds))

	if err := sense.IngestBatcher(context.Background(), ds); err != nil {
		log.Fatalf("fail: IngestBatcher: %s", err)
	}

	log.Println("Serving Mops")
	grpcServer.Serve(listen)
}

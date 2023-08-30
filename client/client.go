package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tinkershack/muxfarm/plumber"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:50050", opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewMopsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	min := pb.NewMediaIn()
	min.Add(pb.StorageType_STORAGE_S3,
		&pb.URI_S3Path{
			S3Path: "https://aomtest1.s3.ap-south-1.amazonaws.com/glocken.mov",
		})
	min.Cid = &pb.CallerID{Cid: "tc2"}
	mid, err := client.Ingest(ctx, min)
	if err != nil {
		log.Fatalf("Fail client.Ingest : %v", err)
	}

	log.Println("return client.Ingest: ", mid)

}

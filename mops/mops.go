package mops

import (
	"log"

	pb "github.com/tinkershack/muxfarm/plumber"
)

type Server struct {
	pb.UnimplementedMopsServer
	// mu sync.Mutex
}

func (s *Server) Mimo(mediaIn *pb.MediaIn, stream pb.Mops_MimoServer) error {
	log.Println("MimoServer")
	log.Printf("MediaIn: %+v", mediaIn)

	for i := 0; i < 3; i++ {
		if err := stream.Send(new(pb.MediaOut)); err != nil {
			return err
		}
	}
	return nil
}

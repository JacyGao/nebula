package server

import (
	"context"
	"log"

	pb "nebula/proto"
)

// server is used to implement the nebula server.
type Server struct {
	pb.UnimplementedNebulaServer
}

// PingPong implements the nebula server
func (s *Server) PingPong(ctx context.Context, in *pb.PingPongRequest) (*pb.PingPongReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PingPongReply{Message: "Pong"}, nil
}

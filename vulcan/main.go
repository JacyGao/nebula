// Package main implements a GRPC server
package main

import (
	"log"
	"net"

	pb "nebula/proto"
	"nebula/server"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNebulaServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

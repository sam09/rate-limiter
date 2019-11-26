//go:generate protoc -I ../token-bucket --go_out=plugins=grpc:../token-bucket ../token-bucket/token-bucket.proto

// Package main implements a server for TokenBucket service.
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sam09/rate-limiter/token-bucket"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// rateLimitServer is used to implement TokenBucket.
type rateLimitServer struct {
	pb.UnimplementedTokenBucketServer
}

func (s *rateLimitServer) AddToken(ctx context.Context, in *pb.AddTokenRequest) (*pb.AddTokenResponse, error) {
	log.Printf("Received: %v", in.GetToken())
	return &pb.AddTokenResponse{Done: true}, nil
}

func (s *rateLimitServer) ConsumeToken(ctx context.Context, in *pb.ConsumeTokenRequest) (*pb.ConsumeTokenResponse, error) {
	return &pb.ConsumeTokenResponse{Token: nil}, nil
}

func (s *rateLimitServer) Refill(ctx context.Context, in *pb.RefillTokenRequest) (*pb.RefillTokenResponse, error) {
	return &pb.RefillTokenResponse{Done: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenBucketServer(s, &rateLimitServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
